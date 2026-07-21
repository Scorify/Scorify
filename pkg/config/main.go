package config

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
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

	// Minion is the configuration for the minion
	Minion struct {
		// id is the id of the minion
		ID uuid.UUID
	}

	// RabbitMQ is the configuration for the RabbitMQ server
	RabbitMQ struct {
		// Host is the host of the RabbitMQ server
		Host string

		// Port is the port of the RabbitMQ server
		Port int

		// Server is the configuration for the RabbitMQ server
		Server struct {
			// User is the user of the RabbitMQ server
			User string

			// Password is the password of the RabbitMQ server
			Password string
		}

		// Minion is the configuration for the RabbitMQ server
		Minion struct {
			// User is the user of the RabbitMQ server
			User string

			// Password is the password of the RabbitMQ server
			Password string

			// QoS is the prefetch count for minion task consumers
			QoS int
		}
	}
)

func InitMinion() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Error("failed to load .env file")
	}

	domain()
	port()
	interval()
	minionID()
	rabbitmqMinion()
}

func InitKoth() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Error("failed to load .env file")
	}

	domain()
	port()
	interval()
	minionID()
	rabbitmqServer()
}

func InitServer() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Error("failed to load .env file")
	}

	domain()
	port()
	interval()
	jwt()
	postgres()
	redis()
	rabbitmqServer()
}

func domain() {
	Domain = os.Getenv("DOMAIN")
	if Domain == "" {
		logrus.Fatal("DOMAIN is not set")
	}
}

func port() {
	var err error

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse PORT")
	}
}

func interval() {
	var err error

	IntervalStr = os.Getenv("INTERVAL")
	Interval, err = time.ParseDuration(IntervalStr)
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse INTERVAL")
	}
	if Interval < 5*time.Second {
		logrus.Fatal("INTERVAL must be greater than 5 second")
	}
}

func jwt() {
	var err error

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
}

func postgres() {
	var err error

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
}

func redis() {
	var err error

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
}

func minionID() {
	_, err := os.Stat(".minion")
	if os.IsNotExist(err) {
		// .minion file does not exist

		minion_id_string := os.Getenv("MINION_ID")
		if minion_id_string == "" {
			Minion.ID = uuid.New()
		} else {
			Minion.ID, err = uuid.Parse(minion_id_string)
			if err != nil {
				logrus.WithError(err).Fatal("failed to parse MINION_ID")
			}
		}

		file, err := os.Create(".minion")
		if err != nil {
			logrus.WithError(err).Fatal("failed to create .minion file")
		}

		_, err = file.WriteString(Minion.ID.String())
		if err != nil {
			logrus.WithError(err).Fatal("failed to write to .minion file")
		}
	} else if err == nil {
		file, err := os.Open(".minion")
		if err != nil {
			logrus.WithError(err).Fatal("failed to open .minion file")
		}

		var out bytes.Buffer
		_, err = io.Copy(&out, file)
		if err != nil {
			logrus.WithError(err).Fatal("failed to read .minion file")
		}

		Minion.ID, err = uuid.Parse(out.String())
		if err != nil {
			logrus.WithError(err).Fatal("failed to parse minion id from .minion file")
		}
	} else {
		logrus.WithError(err).Fatal("failed to open .minion file")
	}
}

func rabbitmqServer() {
	var err error

	RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	if RabbitMQ.Host == "" {
		logrus.Fatal("RABBITMQ_HOST is not set")
	}

	RabbitMQ.Port, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse RABBITMQ_PORT")
	}

	RabbitMQ.Server.User = os.Getenv("RABBITMQ_DEFAULT_USER")
	if RabbitMQ.Server.User == "" {
		logrus.Fatal("RABBITMQ_DEFAULT_USER is not set")
	}

	RabbitMQ.Server.Password = os.Getenv("RABBITMQ_DEFAULT_PASS")
	if RabbitMQ.Server.Password == "" {
		logrus.Fatal("RABBITMQ_DEFAULT_PASS is not set")
	}

	RabbitMQ.Minion.User = os.Getenv("RABBITMQ_MINION_USER")
	if RabbitMQ.Minion.User == "" {
		logrus.Fatal("RABBITMQ_MINION_USER is not set")
	}

	RabbitMQ.Minion.Password = os.Getenv("RABBITMQ_MINION_PASS")
	if RabbitMQ.Minion.Password == "" {
		logrus.Fatal("RABBITMQ_MINION_PASS is not set")
	}
}

func rabbitmqMinion() {
	var err error

	RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	if RabbitMQ.Host == "" {
		logrus.Fatal("RABBITMQ_HOST is not set")
	}

	RabbitMQ.Port, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse RABBITMQ_PORT")
	}

	RabbitMQ.Minion.User = os.Getenv("RABBITMQ_MINION_USER")
	if RabbitMQ.Minion.User == "" {
		logrus.Fatal("RABBITMQ_MINION_USER is not set")
	}

	RabbitMQ.Minion.Password = os.Getenv("RABBITMQ_MINION_PASS")
	if RabbitMQ.Minion.Password == "" {
		logrus.Fatal("RABBITMQ_MINION_PASS is not set")
	}

	RabbitMQ.Minion.QoS, err = strconv.Atoi(os.Getenv("RABBITMQ_MINION_QOS"))
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse RABBITMQ_MINION_QOS")
	}
	if RabbitMQ.Minion.QoS < 1 {
		logrus.Fatal("RABBITMQ_MINION_QOS must be at least 1")
	}
}
