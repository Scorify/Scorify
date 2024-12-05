package auth

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/structs"
)

func JWTMiddleware(entClient *ent.Client, redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("auth")
		if err != nil {
			ctx.Next()
			return
		}

		ok := cache.GetAuth(ctx, redisClient, tokenString)
		if !ok {
			ctx.Next()
			return
		}

		claims := &structs.Claims{}
		jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT.Secret), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.Next()
				return
			}
			ctx.Next()
			return
		}
		if !jwtToken.Valid {
			ctx.Next()
			return
		}

		if claims.Become != "" {
			// Add Become User to Context
			becomeUUID, err := uuid.Parse(claims.Become)
			if err != nil {
				ctx.Next()
				return
			}

			entUser, err := entClient.User.
				Query().
				Where(
					user.IDEQ(becomeUUID),
				).
				Only(ctx.Request.Context())
			if err != nil {
				ctx.Next()
				return
			}

			ctx.Request = ctx.Request.WithContext(
				context.WithValue(
					ctx.Request.Context(),
					structs.USER_CTX_KEY,
					entUser,
				),
			)

			// Add Admin User to Context
			uuid, err := uuid.Parse(claims.ID)
			if err != nil {
				ctx.Next()
				return
			}

			entUser, err = entClient.User.
				Query().
				Where(
					user.IDEQ(uuid),
				).
				Only(ctx.Request.Context())
			if err != nil {
				ctx.Next()
				return
			}

			ctx.Request = ctx.Request.WithContext(
				context.WithValue(
					ctx.Request.Context(),
					structs.ADMIN_CTX_KEY,
					entUser,
				),
			)
		} else {
			uuid, err := uuid.Parse(claims.ID)
			if err != nil {
				ctx.Next()
				return
			}

			entUser, err := entClient.User.
				Query().
				Where(
					user.IDEQ(uuid),
				).
				Only(ctx.Request.Context())
			if err != nil {
				ctx.Next()
				return
			}

			ctx.Request = ctx.Request.WithContext(
				context.WithValue(
					ctx.Request.Context(),
					structs.USER_CTX_KEY,
					entUser,
				),
			)
		}

		ctx.Request = ctx.Request.WithContext(
			context.WithValue(
				ctx.Request.Context(),
				structs.CLAIMS_CTX_KEY,
				claims,
			),
		)

		ctx.Request = ctx.Request.WithContext(
			context.WithValue(
				ctx.Request.Context(),
				structs.TOKEN_CTX_KEY,
				tokenString,
			),
		)

		ctx.Request = ctx.Request.WithContext(
			context.WithValue(
				ctx.Request.Context(),
				structs.CLIENT_IP_CTX_KEY,
				ctx.ClientIP(),
			),
		)

		ctx.Next()
	}
}

func Parse(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(structs.USER_CTX_KEY).(*ent.User)
	if !ok {
		return nil, fmt.Errorf("invalid user")
	}
	return user, nil
}

func ParseAdmin(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(structs.ADMIN_CTX_KEY).(*ent.User)
	if !ok {
		return nil, fmt.Errorf("invalid user")
	}
	return user, nil
}

func ParseClaims(ctx context.Context) (*structs.Claims, error) {
	claims, ok := ctx.Value(structs.CLAIMS_CTX_KEY).(*structs.Claims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
}

func ParseToken(ctx context.Context) (string, error) {
	token, ok := ctx.Value(structs.TOKEN_CTX_KEY).(string)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	return token, nil
}

func ParseClientIP(ctx context.Context) (string, error) {
	clientIP, ok := ctx.Value(structs.CLIENT_IP_CTX_KEY).(string)
	if !ok {
		return "", fmt.Errorf("invalid ip")
	}
	return clientIP, nil
}
