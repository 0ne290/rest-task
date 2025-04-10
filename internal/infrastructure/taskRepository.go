package infrastructure

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"rest-task/internal/core/entities"
)

type PosgresTaskRepository struct {
	transaction pgx.Tx
}

func newPosgresTaskRepository(transaction pgx.Tx) *PosgresTaskRepository {
	return &PosgresTaskRepository{transaction}
}

func (r *PosgresTaskRepository) Create(ctx context.Context, task *entities.Task) error {
	const query string = "INSERT INTO tasks VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.transaction.Exec(ctx, query, task.Uuid, task.UserUuid, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)

	return err
}

func (r *PosgresTaskRepository) GetAllByUser(ctx context.Context, userUuid uuid.UUID) ([]*entities.Task, error) {
	const query string = "SELECT * FROM tasks WHERE user_uuid = $1"

	rows, err := r.transaction.Query(ctx, query, userUuid)
	if err != nil {
		return make([]*entities.Task, 0), err
	}
	defer rows.Close()

	var tasks []*entities.Task
	for rows.Next() {
		task := &entities.Task{}

		err = rows.Scan(&task.Uuid, &task.UserUuid, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return make([]*entities.Task, 0), err
		}

		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return make([]*entities.Task, 0), err
	}

	return tasks, nil
}

func (r *PosgresTaskRepository) TryGetByUuid(ctx context.Context, taskUuid uuid.UUID) (*entities.Task, error) {
	const query string = "SELECT * FROM tasks WHERE uuid = $1 FOR UPDATE"

	task := &entities.Task{}

	err := r.transaction.QueryRow(ctx, query, taskUuid).Scan(&task.Uuid, &task.UserUuid, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return task, nil
}

func (r *PosgresTaskRepository) Update(ctx context.Context, task *entities.Task) error {
	const query string = "UPDATE tasks SET status = $2, updated_at = $3 WHERE uuid = $1"

	_, err := r.transaction.Exec(ctx, query, task.Uuid, task.Status, task.UpdatedAt)

	return err
}

func (r *PosgresTaskRepository) TryDeleteByUuid(ctx context.Context, taskUuid uuid.UUID) (*entities.Task, error) {
	const query string = "DELETE FROM tasks WHERE uuid = $1 RETURNING *"

	task := &entities.Task{}

	err := r.transaction.QueryRow(ctx, query, taskUuid).Scan(&task.Uuid, &task.UserUuid, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return task, nil
}
