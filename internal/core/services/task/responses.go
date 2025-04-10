package taskService

import (
	"github.com/google/uuid"
	"rest-task/internal/core/entities"
)

type CreateResponse struct {
	Uuid uuid.UUID `json:"uuid"`
}

type GetAllResponse struct {
	Tasks []*entities.TaskView `json:"tasks"`
}

type GetByUuidResponse struct {
	Task *entities.TaskView `json:"task"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type UpdateResponse struct {
	Status entities.Status `json:"status"`
}
