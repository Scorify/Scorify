package data

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/helpers"
	"github.com/sirupsen/logrus"
)

func NewClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Postgres.Host,
			config.Postgres.Port,
			config.Postgres.User,
			config.Postgres.Password,
			config.Postgres.DB,
		),
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed opening connection to postgres")
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		logrus.WithError(err).Fatalf("failed creating schema resources")
		return nil, err
	}

	exists, err := client.User.Query().
		Where(
			user.UsernameEQ("admin"),
		).Exist(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("failed checking if admin user exists")
		return nil, err
	}

	if !exists {
		hashedPassword, err := helpers.HashPassword("admin")
		if err != nil {
			logrus.WithError(err).Fatalf("failed hashing admin password")
			return nil, err
		}

		_, err = client.User.Create().
			SetUsername("admin").
			SetPassword(hashedPassword).
			SetRole(user.RoleAdmin).
			Save(ctx)
		if err != nil {
			logrus.WithError(err).Warnf("failed creating admin user")
			return nil, err
		}
	}

	return client, nil
}
