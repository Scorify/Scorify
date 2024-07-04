package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	// Domain is the domain of the cookie
	Domain string

	// Port is the port of the server
	Port int

	// Scoring Interval is the interval for the scoring in seconds
	IntervalStr string
	Interval    time.Duration

	// JWT is the configuration for the JWT token
	JWT struct {
		// Timeout is the timeout for the JWT token in hours
		TimeoutStr string
		Timeout    time.Duration

		// Key is the secret key for the JWT token
		Secret string
	}

	// Postgres is the configuration for the postgres database
	Postgres struct {
		// Host is the host of the postgres database
		Host string

		// Port is the port of the postgres database
		Port int

		// User is the user of the postgres database
		User string

		// Password is the password of the postgres database
		Password string

		// DB is the name of the postgres database
		DB string
	}

	// Redis is the configuration for the redis server
	Redis struct {
		// Host is the host of the redis server
		Host string

		// Port is the port of the redis server
		Port int

		// Password is the password of the redis server
		Password string
	}

	// gRPC is the configuration for the gRPC server
	GRPC struct {
		// Host is the host of the gRPC server
		Host string

		// Port is the port of the gRPC server
		Port int

		// Secret is the secret key for the gRPC server
		Secret string
	}
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("failed to load .env file")
	}

	Domain = os.Getenv("DOMAIN")
	if Domain == "" {
		logrus.Fatal("DOMAIN is not set")
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse PORT")
	}

	IntervalStr = os.Getenv("INTERVAL")
	Interval, err = time.ParseDuration(IntervalStr)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse INTERVAL")
	}
	if Interval < time.Second {
		logrus.Fatal("INTERVAL must be greater than 1 second")
	}

	JWT.TimeoutStr = os.Getenv("JWT_TIMEOUT")
	JWT.Timeout, err = time.ParseDuration(JWT.TimeoutStr)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse JWT_TIMEOUT")
	}
	if JWT.Timeout <= time.Hour {
		logrus.Fatal("JWT_TIMEOUT must be greater than 1 hour")
	}

	JWT.Secret = os.Getenv("JWT_SECRET")
	if JWT.Secret == "" {
		logrus.Fatal("JWT_SECRET is not set")
	}

	Postgres.Host = os.Getenv("POSTGRES_HOST")
	if Postgres.Host == "" {
		logrus.Fatal("POSTGRES_HOST is not set")
	}

	Postgres.Port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse POSTGRES_PORT")
	}

	Postgres.User = os.Getenv("POSTGRES_USER")
	if Postgres.User == "" {
		logrus.Fatal("POSTGRES_USER is not set")
	}

	Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	if Postgres.Password == "" {
		logrus.Fatal("POSTGRES_PASSWORD is not set")
	}

	Postgres.DB = os.Getenv("POSTGRES_DB")
	if Postgres.DB == "" {
		logrus.Fatal("POSTGRES_DB is not set")
	}

	Redis.Host = os.Getenv("REDIS_HOST")
	if Redis.Host == "" {
		logrus.Fatal("REDIS_HOST is not set")
	}

	Redis.Port, err = strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse REDIS_PORT")
	}

	Redis.Password = os.Getenv("REDIS_PASSWORD")
	if Redis.Password == "" {
		logrus.Fatal("REDIS_PASSWORD is not set")
	}

	GRPC.Host = os.Getenv("GRPC_HOST")
	if GRPC.Host == "" {
		logrus.Fatal("GRPC_HOST is not set")
	}

	GRPC.Port, err = strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse GRPC_PORT")
	}

	GRPC.Secret = os.Getenv("GRPC_SECRET")
	if GRPC.Secret == "" {
		logrus.Fatal("GRPC_SECRET is not set")
	}
}
