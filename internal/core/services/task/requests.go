package taskService

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	UserUuid    uuid.UUID `json:"userUuid"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
}

type GetAllRequest struct {
	UserUuid uuid.UUID `json:"userUuid"`
}

type GetByUuidRequest struct {
	UserUuid uuid.UUID `json:"userUuid"`
	TaskUuid uuid.UUID `json:"taskUuid"`
}

type DeleteRequest struct {
	UserUuid uuid.UUID `json:"userUuid"`
	TaskUuid uuid.UUID `json:"taskUuid"`
}

type UpdateRequest struct {
	UserUuid uuid.UUID `json:"userUuid"`
	TaskUuid uuid.UUID `json:"taskUuid"`
}
