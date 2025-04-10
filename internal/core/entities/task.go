package entities

import (
	"github.com/google/uuid"
	"rest-task/internal/core"
	"time"
)

type Status string

const (
	statusNew        Status = "new"
	statusInProgress Status = "in_progress"
	statusDone       Status = "done"
)

type TaskView struct {
	Uuid        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Task struct {
	Uuid        uuid.UUID
	UserUuid    uuid.UUID
	Title       string
	Description *string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(uuid uuid.UUID, userUuid uuid.UUID, title string, description *string, createdAt time.Time, updatedAt time.Time) *Task {
	return &Task{uuid, userUuid, title, description, statusNew, createdAt, updatedAt}
}

func (task *Task) ToView() *TaskView {
	return &TaskView{task.Uuid, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt}
}

func (task *Task) Update(updatedAt time.Time) error {
	switch task.Status {

	case statusNew:
		task.Status = statusInProgress
		task.UpdatedAt = updatedAt

		return nil

	case statusInProgress:
		task.Status = statusDone
		task.UpdatedAt = updatedAt

		return nil

	default:
		return &core.InvariantViolationError{Message: "status is invalid"}
	}
}
