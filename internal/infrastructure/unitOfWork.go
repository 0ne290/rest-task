package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v5"
	"rest-task/internal/core/services"
)

type postgresUnitOfWork struct {
	transaction    pgx.Tx
	taskRepository *PosgresTaskRepository
}

func newPostgresUnitOfWork(transaction pgx.Tx) *postgresUnitOfWork {
	return &postgresUnitOfWork{transaction, newPosgresTaskRepository(transaction)}
}

func (uow *postgresUnitOfWork) TaskRepository() services.TaskRepository {
	return uow.taskRepository
}

func (uow *postgresUnitOfWork) Save(ctx context.Context) error {
	return uow.transaction.Commit(ctx)
}

func (uow *postgresUnitOfWork) Rollback(ctx context.Context) error {
	return uow.transaction.Rollback(ctx)
}
