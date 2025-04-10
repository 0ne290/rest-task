package taskWeb

import (
	"context"
	"rest-task/internal/core/services/task"
)

type Service interface {
	Create(ctx context.Context, request *taskService.CreateRequest) (*taskService.CreateResponse, error)
	GetAll(ctx context.Context, request *taskService.GetAllRequest) (*taskService.GetAllResponse, error)
	GetByUuid(ctx context.Context, request *taskService.GetByUuidRequest) (*taskService.GetByUuidResponse, error)
	Update(ctx context.Context, request *taskService.UpdateRequest) (*taskService.UpdateResponse, error)
	Delete(ctx context.Context, request *taskService.DeleteRequest) (*taskService.DeleteResponse, error)
}
