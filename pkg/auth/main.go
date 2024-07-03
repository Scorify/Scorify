package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/structs"
)

func GenerateJWT(username string, id uuid.UUID, become *uuid.UUID) (string, int, error) {
	expiration := time.Now().Add(config.JWT.Timeout)

	claims := &structs.Claims{
		ID: id.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	if become != nil {
		claims.Become = become.String()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.JWT.Secret))

	return tokenStr, int(expiration.Unix()), err
}

func ParseJWT(tokenString string) (*jwt.Token, *structs.Claims, error) {
	claims := &structs.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWT.Secret), nil
	})
	return token, claims, err
}
