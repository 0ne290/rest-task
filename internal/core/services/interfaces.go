package services

import (
	"context"
	"github.com/google/uuid"
	"rest-task/internal/core/entities"
	"time"
)

type UnitOfWorkStarter interface {
	Start(ctx context.Context) (UnitOfWork, error)
}

type UnitOfWork interface {
	TaskRepository() TaskRepository

	Save(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) error
	GetAllByUser(ctx context.Context, userUuid uuid.UUID) ([]*entities.Task, error)
	TryGetByUuid(ctx context.Context, taskUuid uuid.UUID) (*entities.Task, error)
	Update(ctx context.Context, task *entities.Task) error
	TryDeleteByUuid(ctx context.Context, taskUuid uuid.UUID) (*entities.Task, error)
}

type TimeProvider interface {
	Now() time.Time
}

type UuidProvider interface {
	Random() uuid.UUID
}
