package taskWeb

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rest-task/internal/core/services/task"
	"rest-task/internal/web"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

// @Summary Create a new user task with status "new". JWT authentication with claim "userUuid" is required
// @Accept json
// @Produce json
// @Param createRequest body CreateRequest true "CreateRequest"
// @Success 200 {object} taskService.CreateResponse
// @Failure 400 {object} web.Error
// @Failure 401 {object} web.Error
// @Failure 500 {object} web.Error500
// @Router /v1/tasks [post]
func (c *Controller) Create(ctx *fiber.Ctx) error {
	request := &CreateRequest{}
	err := ctx.BodyParser(request)
	if err != nil {
		return web.Create400(ctx, &web.Error{Message: "invalid request body format"})
	}

	response, err := c.service.Create(ctx.Context(), mapCreateRequest(request, ctx.Locals("userUuid").(uuid.UUID)))
	if err != nil {
		return err
	}

	return web.Create200(ctx, response)
}

func mapCreateRequest(source *CreateRequest, userUuid uuid.UUID) *taskService.CreateRequest {
	return &taskService.CreateRequest{UserUuid: userUuid, Title: source.Title, Description: source.Description}
}

// @Summary Get all user tasks. JWT authentication with claim "userUuid" is required
// @Accept json
// @Produce json
// @Success 200 {object} taskService.GetAllResponse
// @Failure 400 {object} web.Error
// @Failure 401 {object} web.Error
// @Failure 500 {object} web.Error500
// @Router /v1/tasks [get]
func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	response, err := c.service.GetAll(ctx.Context(), &taskService.GetAllRequest{UserUuid: ctx.Locals("userUuid").(uuid.UUID)})
	if err != nil {
		return err
	}

	return web.Create200(ctx, response)
}

// @Summary Get user task by ID. JWT authentication with claim "userUuid" is required
// @Produce json
// @Success 200 {object} taskService.GetByUuidResponse
// @Failure 400 {object} web.Error
// @Failure 401 {object} web.Error
// @Failure 404 {object} web.Error
// @Failure 500 {object} web.Error500
// @Param uuid path string true "uuid" Format(uuid)
// @Router /v1/tasks/{uuid} [get]
func (c *Controller) GetByUuid(ctx *fiber.Ctx) error {
	taskUuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return web.Create400(ctx, &web.Error{Message: "invalid request body format"})
	}

	response, err := c.service.GetByUuid(ctx.Context(), &taskService.GetByUuidRequest{UserUuid: ctx.Locals("userUuid").(uuid.UUID), TaskUuid: taskUuid})
	if err != nil {
		return err
	}

	return web.Create200(ctx, response)
}

// @Summary Delete a user task by UUID. JWT authentication with claim "userUuid" is required
// @Produce json
// @Success 200 {object} taskService.DeleteResponse
// @Failure 400 {object} web.Error
// @Failure 401 {object} web.Error
// @Failure 404 {object} web.Error
// @Failure 500 {object} web.Error500
// @Param uuid path string true "uuid" Format(uuid)
// @Router /v1/tasks/{uuid} [delete]
func (c *Controller) Delete(ctx *fiber.Ctx) error {
	taskUuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return web.Create400(ctx, &web.Error{Message: "invalid request body format"})
	}

	response, err := c.service.Delete(ctx.Context(), &taskService.DeleteRequest{UserUuid: ctx.Locals("userUuid").(uuid.UUID), TaskUuid: taskUuid})
	if err != nil {
		return err
	}

	return web.Create200(ctx, response)
}

// @Summary Moves the user task to the next status. JWT authentication with claim "userUuid" is required
// @Description Transitions of state machine: new -> in_progress -> done
// @Produce json
// @Success 200 {object} taskService.UpdateResponse
// @Failure 400 {object} web.Error
// @Failure 401 {object} web.Error
// @Failure 404 {object} web.Error
// @Failure 500 {object} web.Error500
// @Param uuid path string true "uuid" Format(uuid)
// @Router /v1/tasks/{uuid} [put]
func (c *Controller) Update(ctx *fiber.Ctx) error {
	taskUuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return web.Create400(ctx, &web.Error{Message: "invalid request body format"})
	}

	response, err := c.service.Update(ctx.Context(), &taskService.UpdateRequest{UserUuid: ctx.Locals("userUuid").(uuid.UUID), TaskUuid: taskUuid})
	if err != nil {
		return err
	}

	return web.Create200(ctx, response)
}
