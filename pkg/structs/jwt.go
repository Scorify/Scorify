package structs

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID     string `json:"id"`
	Become string `json:"become"`
	jwt.RegisteredClaims
}

type contextKey struct {
	name string
}

var (
	USER_CTX_KEY  = &contextKey{"user"}
	ADMIN_CTX_KEY = &contextKey{"admin"}
)
