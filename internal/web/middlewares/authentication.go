package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strings"
)

func Authentication(jwtManager JwtManager) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// проверка токена авторизации
		splittedAuthHeader := strings.Split(ctx.Get("Authorization"), "Bearer ")
		if len(splittedAuthHeader) == 2 {
			userUuid := jwtManager.ParseUserUuid(splittedAuthHeader[1])
			if userUuid == uuid.Nil {
				return ctx.SendStatus(fiber.StatusUnauthorized)
			}

			ctx.Locals("userUuid", userUuid)

			return ctx.Next()
		}

		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
}
