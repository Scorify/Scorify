package setup

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "setup",
	Short:   "Setup configuration for the server",
	Long:    "Setup configuration for the server",
	Aliases: []string{"init", "i"},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	// Create .env file if it doesn't exist
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		fmt.Println("[X] .env file not found")
		fmt.Println("[*] Creating .env file")
		fmt.Println()
		err = createMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show create menu")
		}
		return
	} else if err != nil {
		logrus.WithError(err).Fatal("failed to check .env file")
	}

	choice, err := actionMenu()
	if err != nil {
		logrus.WithError(err).Fatal("failed to show action menu")
	}

	switch choice {
	case actionCreate:
		err = createMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show create menu")
		}
	case actionUpdate:
		config.InitServer()
		err = editMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show edit menu")
		}
	case actionDelete:
		err = deleteMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show delete menu")
		}
	case actionView:
		config.InitServer()
		err = viewMenu()
		if err != nil {
			logrus.WithError(err).Fatal("failed to show view menu")
		}
	case actionNone:
		return
	}
}

func createMenu() error {
	reader := bufio.NewReader(os.Stdin)

	// DOMAIN
	domain, err := prompt(
		reader,
		"localhost",
		"Enter the domain of the server [localhost]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read domain: %w", err)
	}

	// PORT
	port, err := promptInt(
		reader,
		8080,
		"Enter the port of the server [8080]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read port: %w", err)
	}

	// INTERVAL
	interval, err := promptDuration(
		reader,
		30*time.Second,
		time.Second,
		"Enter the interval of the score task in seconds [30s]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read interval: %w", err)
	}

	// JWT_TIMEOUT
	jwtTimeout, err := promptDuration(
		reader,
		6*time.Hour,
		time.Hour,
		"Enter the timeout of the JWT (session length) in hours [6h]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read JWT timeout: %w", err)
	}

	// JWT_SECRET
	jwtSecret, err := promptPassword(
		reader,
		"Enter the secret key for the JWT token [randomly generate]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read JWT secret: %w", err)
	}

	// POSTGRES_HOST
	postgresHost, err := prompt(
		reader,
		"postgres",
		"Enter the host of the postgres database [postgres]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read postgres host: %w", err)
	}

	// POSTGRES_PORT
	postgresPort, err := promptInt(
		reader,
		5432,
		"Enter the port of the postgres database [5432]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read postgres port: %w", err)
	}

	// POSTGRES_USER
	postgresUser, err := prompt(
		reader,
		"scorify",
		"Enter the user of the postgres database [scorify]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read postgres user: %w", err)
	}

	// POSTGRES_PASSWORD
	postgresPassword, err := promptPassword(
		reader,
		"Enter the password of the postgres database [randomly generate]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read postgres password: %w", err)
	}

	// POSTGRES_DB
	postgresDB, err := prompt(
		reader,
		"scorify",
		"Enter the name of the postgres database [scorify]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read postgres database: %w", err)
	}

	// REDIS_HOST
	redisHost, err := prompt(
		reader,
		"redis",
		"Enter the host of the redis server [redis]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read redis host: %w", err)
	}

	// REDIS_PORT
	redisPort, err := promptInt(
		reader,
		6379,
		"Enter the port of the redis server [6379]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read redis port: %w", err)
	}

	// REDIS_PASSWORD
	redisPassword, err := promptPassword(
		reader,
		"Enter the password of the redis server [randomly generate]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read redis password: %w", err)
	}

	// GRPC_HOST
	grpcHost, err := prompt(
		reader,
		"localhost",
		"Enter the host of the gRPC server [localhost]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read gRPC host: %w", err)
	}

	// GRPC_PORT
	grpcPort, err := promptInt(
		reader,
		50051,
		"Enter the port of the gRPC server [50051]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read gRPC port: %w", err)
	}

	// GRPC_SECRET
	grpcSecret, err := promptPassword(
		reader,
		"Enter the secret key for the gRPC server [randomly generate]: ",
	)
	if err != nil {
		return fmt.Errorf("failed to read gRPC secret: %w", err)
	}

	err = writeConfig(
		domain,
		port,
		interval,
		jwtTimeout,
		jwtSecret,
		postgresHost,
		postgresPort,
		postgresUser,
		postgresPassword,
		postgresDB,
		redisHost,
		redisPort,
		redisPassword,
		grpcHost,
		grpcPort,
		grpcSecret,
	)
	if err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

func writeConfig(
	domain string,
	port int,
	interval time.Duration,
	jwtTimeout time.Duration,
	jwtSecret string,
	postgresHost string,
	postgresPort int,
	postgresUser string,
	postgresPassword string,
	postgresDB string,
	redisHost string,
	redisPort int,
	redisPassword string,
	grpcHost string,
	grpcPort int,
	grpcSecret string,
) error {
	envTmpl, err := os.ReadFile(".env.tmpl")
	if err != nil {
		return fmt.Errorf("failed to read .env.tmpl: %w", err)
	}

	tmpl, err := template.New("env").Parse(string(envTmpl))
	if err != nil {
		return fmt.Errorf("failed to parse .env.tmpl: %w", err)
	}

	envFile, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("failed to create .env: %w", err)
	}

	err = tmpl.Execute(envFile, struct {
		Domain     string
		Port       int
		Interval   time.Duration
		JWTTimeout time.Duration
		JWTSecret  string

		PostgresHost     string
		PostgresPort     int
		PostgresUser     string
		PostgresPassword string
		PostgresDB       string

		RedisHost     string
		RedisPort     int
		RedisPassword string

		GRPCHost   string
		GRPCPort   int
		GRPCSecret string
	}{
		Domain:     domain,
		Port:       port,
		Interval:   interval,
		JWTTimeout: jwtTimeout,
		JWTSecret:  jwtSecret,

		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDB:       postgresDB,

		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,

		GRPCHost:   grpcHost,
		GRPCPort:   grpcPort,
		GRPCSecret: grpcSecret,
	})
	if err != nil {
		return fmt.Errorf("failed to execute .env.tmpl: %w", err)
	}

	return nil
}
