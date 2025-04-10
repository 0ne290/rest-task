package taskService

import (
	"context"
	"rest-task/internal/core"
	"rest-task/internal/core/entities"
	"rest-task/internal/core/services"
)

type RealService struct {
	unitOfWorkStarter services.UnitOfWorkStarter
	timeProvider      services.TimeProvider
	uuidProvider      services.UuidProvider
}

func NewRealService(unitOfWorkStarter services.UnitOfWorkStarter, timeProvider services.TimeProvider, uuidProvider services.UuidProvider) *RealService {
	return &RealService{unitOfWorkStarter, timeProvider, uuidProvider}
}

func (s *RealService) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	unitOfWork, err := s.unitOfWorkStarter.Start(ctx)
	if err != nil {
		return nil, err
	}
	taskRepository := unitOfWork.TaskRepository()

	timeNow := s.timeProvider.Now()
	task := entities.NewTask(s.uuidProvider.Random(), request.UserUuid, request.Title, request.Description, timeNow, timeNow)

	err = taskRepository.Create(ctx, task)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	err = unitOfWork.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &CreateResponse{task.Uuid}, nil
}

func (s *RealService) GetAll(ctx context.Context, request *GetAllRequest) (*GetAllResponse, error) {
	unitOfWork, err := s.unitOfWorkStarter.Start(ctx)
	if err != nil {
		return nil, err
	}
	taskRepository := unitOfWork.TaskRepository()

	tasks, err := taskRepository.GetAllByUser(ctx, request.UserUuid)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	err = unitOfWork.Save(ctx)
	if err != nil {
		return nil, err
	}

	tasksView := make([]*entities.TaskView, 0, len(tasks))
	for _, task := range tasks {
		tasksView = append(tasksView, task.ToView())
	}

	return &GetAllResponse{tasksView}, nil
}

func (s *RealService) GetByUuid(ctx context.Context, request *GetByUuidRequest) (*GetByUuidResponse, error) {
	unitOfWork, err := s.unitOfWorkStarter.Start(ctx)
	if err != nil {
		return nil, err
	}
	taskRepository := unitOfWork.TaskRepository()

	task, err := taskRepository.TryGetByUuid(ctx, request.TaskUuid)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	if task == nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.NotFoundError{Message: "task does not exists"}
	}
	if task.UserUuid != request.UserUuid {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.PermissionDeniedError{Message: "permission denied"}
	}

	err = unitOfWork.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &GetByUuidResponse{task.ToView()}, nil
}

func (s *RealService) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, error) {
	unitOfWork, err := s.unitOfWorkStarter.Start(ctx)
	if err != nil {
		return nil, err
	}
	taskRepository := unitOfWork.TaskRepository()

	task, err := taskRepository.TryGetByUuid(ctx, request.TaskUuid)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	if task == nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.NotFoundError{Message: "task does not exists"}
	}
	if task.UserUuid != request.UserUuid {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.PermissionDeniedError{Message: "permission denied"}
	}

	err = task.Update(s.timeProvider.Now())
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	err = taskRepository.Update(ctx, task)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	err = unitOfWork.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &UpdateResponse{task.Status}, nil
}

func (s *RealService) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	unitOfWork, err := s.unitOfWorkStarter.Start(ctx)
	if err != nil {
		return nil, err
	}
	taskRepository := unitOfWork.TaskRepository()

	deletedTask, err := taskRepository.TryDeleteByUuid(ctx, request.TaskUuid)
	if err != nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, err
	}

	if deletedTask == nil {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.NotFoundError{Message: "task does not exists"}
	}
	if deletedTask.UserUuid != request.UserUuid {
		_ = unitOfWork.Rollback(ctx)

		return nil, &core.PermissionDeniedError{Message: "permission denied"}
	}

	err = unitOfWork.Save(ctx)
	if err != nil {
		return nil, err
	}

	return &DeleteResponse{"task deleted"}, nil
}
