package infrastructure

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RealJwtManager struct {
	key []byte
}

func NewRealJwtManager(key []byte) *RealJwtManager {
	return &RealJwtManager{key}
}

func (jm *RealJwtManager) ParseUserUuid(tokenString string) uuid.UUID {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return jm.key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))
	if err != nil || !token.Valid {
		return uuid.Nil
	}

	claims := token.Claims.(jwt.MapClaims)
	userUuidAny, ok := claims["userUuid"]
	if !ok {
		return uuid.Nil
	}

	userUuidString, ok := userUuidAny.(string)
	if !ok {
		return uuid.Nil
	}

	userUuid, err := uuid.Parse(userUuidString)
	if err != nil {
		return uuid.Nil
	}

	return userUuid
}
