package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"rest-task/internal/core"
	"rest-task/internal/web"
)

var (
	invariantViolationError *core.InvariantViolationError
	permissionDeniedError   *core.PermissionDeniedError
	notFoundError           *core.NotFoundError
)

func ErrorHandlingAndLogging(logger *zap.SugaredLogger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		requestUrl := "[" + ctx.Method() + "] " + ctx.OriginalURL()
		requestUuid := uuid.New().String()

		logger.Infow("begin", "url", requestUrl, "requestUuid", requestUuid, "requestBody", string(ctx.Request().Body()))

		var err error
		ret := ctx.Next()
		if ret != nil {
			if errors.As(ret, &invariantViolationError) {
				ret = web.Create400(ctx, &web.Error{Message: ret.Error()})
			} else if errors.As(ret, &notFoundError) {
				ret = web.Create404(ctx, &web.Error{Message: ret.Error()})
			} else if errors.As(ret, &permissionDeniedError) {
				ret = web.Create401(ctx, &web.Error{Message: ret.Error()})
			} else {
				err = ret

				ret = web.Create500(ctx, web.NewError500(requestUrl, requestUuid))
			}
		}

		response := ctx.Response()
		responseBody := response.Body()
		statusCode := response.StatusCode()
		if statusCode != 500 {
			logger.Infow("end", "requestUuid", requestUuid, "statusCode", statusCode, "responseBody", string(responseBody))
		} else {
			logger.Errorw("end", "requestUuid", requestUuid, "statusCode", statusCode, "responseBody", string(responseBody), "error", err)
		}

		return ret
	}
}
