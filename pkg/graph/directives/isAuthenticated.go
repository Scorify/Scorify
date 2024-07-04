package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/scorify/scorify/pkg/auth"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	_, err = auth.Parse(ctx)
	if err != nil {
		return nil, fmt.Errorf("request not authenticated")
	}

	return next(ctx)
}
