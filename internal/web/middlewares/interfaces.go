package middlewares

import "github.com/google/uuid"

type JwtManager interface {
	ParseUserUuid(tokenString string) uuid.UUID
}
