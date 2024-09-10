package directives

import (
	"context"
	"fmt"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/scorify/scorify/pkg/auth"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/static"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles []*user.Role) (res interface{}, err error) {
	entUser, err := auth.Parse(ctx)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		if *role == entUser.Role {
			return next(ctx)
		}
	}

	entUser, err = auth.ParseAdmin(ctx)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		if *role == entUser.Role {
			return next(ctx)
		}
	}

	return nil, fmt.Errorf(
		"invalid permissions; \"%s\" does not have any of the following roles: [\"%s\"]",
		entUser.Username,
		strings.Join(
			static.MapSlice(
				roles,
				func(_ int, role *user.Role) string {
					return string(*role)
				},
			),
			"\", \"",
		),
	)
}
