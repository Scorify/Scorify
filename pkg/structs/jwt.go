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
	USER_CTX_KEY      = &contextKey{"user"}
	ADMIN_CTX_KEY     = &contextKey{"admin"}
	CLAIMS_CTX_KEY    = &contextKey{"claims"}
	TOKEN_CTX_KEY     = &contextKey{"token"}
	CLIENT_IP_CTX_KEY = &contextKey{"client-ip"}
)
