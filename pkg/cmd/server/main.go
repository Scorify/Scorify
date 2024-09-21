package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/auth"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/data"
	"github.com/scorify/scorify/pkg/engine"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/inject"
	"github.com/scorify/scorify/pkg/ent/injectsubmission"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/graph"
	"github.com/scorify/scorify/pkg/graph/directives"
	"github.com/scorify/scorify/pkg/rabbitmq/management"
	"github.com/scorify/scorify/pkg/rabbitmq/management/vhosts"
	"github.com/scorify/scorify/pkg/rabbitmq/rabbitmq"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run the server",
	Long:    "Run the server",
	Aliases: []string{"s", "serve"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitServer()
	},

	Run: run,
}

func graphqlHandler(entClient *ent.Client, taskRequestChan chan *structs.TaskRequest, taskResponseChan chan *structs.TaskResponse, workerStatusChan chan *structs.WorkerStatus, redisClient *redis.Client, engineClient *engine.Client) gin.HandlerFunc {
	conf := graph.Config{
		Resolvers: &graph.Resolver{
			Ent:              entClient,
			Redis:            redisClient,
			Engine:           engineClient,
			TaskRequestChan:  taskRequestChan,
			TaskResponseChan: taskResponseChan,
			WorkerStatusChan: workerStatusChan,
		},
	}

	conf.Directives.IsAuthenticated = directives.IsAuthenticated
	conf.Directives.HasRole = directives.HasRole

	h := handler.New(
		graph.NewExecutableSchema(
			conf,
		),
	)

	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	if gin.IsDebugging() {
		h.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func injectFileHandler(entClient *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		entUser, err := auth.Parse(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		parentID := c.Param("parentID")
		fileID := c.Param("fileID")
		fileName := c.Param("filename")

		parentUUID, err := uuid.Parse(parentID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid parent id"})
			return
		}

		fileUUID, err := uuid.Parse(fileID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid file id"})
			return
		}

		var entInject *ent.Inject

		if entUser.Role == user.RoleAdmin {
			// Admins can access all files
			entInject, err = entClient.Inject.Query().
				Where(
					inject.ID(parentUUID),
				).
				Only(c)
		} else {
			// Users can only access files that are currently active
			now := time.Now()
			entInject, err = entClient.Inject.Query().
				Where(
					inject.ID(parentUUID),
					inject.StartTimeLTE(now),
					inject.EndTimeGTE(now),
				).
				Only(c)
		}
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("inject not found: %s", err.Error())})
			return
		}

		var file *structs.File

		for _, f := range entInject.Files {
			if f.ID == fileUUID && f.Name == fileName {
				file = &structs.File{
					ID:   f.ID,
					Name: f.Name,
				}
				break
			}
		}

		if file == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}

		filePath, err := file.FilePath(structs.FileTypeInject, parentUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get file path: %s", err.Error())})
			return
		}

		c.File(filePath)
	}
}

func injectSubmissionFileHandler(entClient *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		entUser, err := auth.Parse(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		parentID := c.Param("parentID")
		fileID := c.Param("fileID")
		fileName := c.Param("filename")

		parentUUID, err := uuid.Parse(parentID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid parent id"})
			return
		}

		fileUUID, err := uuid.Parse(fileID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid file id"})
			return
		}

		var entInjectSubmission *ent.InjectSubmission

		if entUser.Role == user.RoleAdmin {
			// Admins can access all files
			entInjectSubmission, err = entClient.InjectSubmission.Query().
				Where(
					injectsubmission.ID(parentUUID),
				).
				Only(c)
		} else {
			// Users can only access files their submissions
			entInjectSubmission, err = entClient.InjectSubmission.Query().
				Where(
					injectsubmission.ID(parentUUID),
					injectsubmission.HasUserWith(
						user.ID(entUser.ID),
					),
				).
				Only(c)
		}
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("failed to get inject submission: %s", err.Error())})
			return
		}

		var file *structs.File

		for _, f := range entInjectSubmission.Files {
			if f.ID == fileUUID && f.Name == fileName {
				file = &structs.File{
					ID:   f.ID,
					Name: f.Name,
				}
				break
			}
		}

		if file == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
			return
		}

		filePath, err := file.FilePath(structs.FileTypeSubmission, parentUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get file path: %s", err.Error())})
			return
		}

		c.File(filePath)
	}
}

func startWebServer(wg *sync.WaitGroup, entClient *ent.Client, redisClient *redis.Client, engineClient *engine.Client, taskRequestChan chan *structs.TaskRequest, taskResponseChan chan *structs.TaskResponse, workerStatusChan chan *structs.WorkerStatus) {
	defer wg.Done()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(auth.JWTMiddleware(entClient))

	err := router.SetTrustedProxies(nil)
	if err != nil {
		logrus.WithError(err).Fatal("failed to set trusted proxies")
	}

	cors_urls := []string{
		fmt.Sprintf("http://%s:%d", config.Domain, config.Port),
		fmt.Sprintf("https://%s:%d", config.Domain, config.Port),
		fmt.Sprintf("http://%s:3000", config.Domain),
		fmt.Sprintf("https://%s:3000", config.Domain),
		fmt.Sprintf("http://%s:5173", config.Domain),
		fmt.Sprintf("https://%s:5173", config.Domain),
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     cors_urls,
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	router.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/api/query")))
	router.POST("/api/query", graphqlHandler(entClient, taskRequestChan, taskResponseChan, workerStatusChan, redisClient, engineClient))
	router.GET("/api/query", graphqlHandler(entClient, taskRequestChan, taskResponseChan, workerStatusChan, redisClient, engineClient))
	router.GET("/api/files/inject/:parentID/:fileID/:filename", injectFileHandler(entClient))
	router.GET("/api/files/submission/:parentID/:fileID/:filename", injectSubmissionFileHandler(entClient))

	logrus.Printf("Starting web server on http://%s:%d", config.Domain, config.Port)

	err = router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.WithError(err).Error("failed to start server")
	} else {
		logrus.Info("Server stopped")
	}
}

func startRabbitMQServer(wg *sync.WaitGroup, ctx context.Context, taskRequestChan chan *structs.TaskRequest, taskResponseChan chan *structs.TaskResponse, workerStatusChan chan *structs.WorkerStatus, redisClient *redis.Client, entClient *ent.Client) {
	defer wg.Done()

	// setup rabbitmq
	managementClient, err := management.Client()
	if err != nil {
		logrus.WithError(err).Error("failed to create rabbitmq management client")
		return
	}

	// setup minion user
	errResp, err := managementClient.Users.Put(
		config.RabbitMQ.Minion.User,
		config.RabbitMQ.Minion.Password,
		nil,
	)
	if err != nil {
		logrus.WithError(err).Error("failed to setup user")
		return
	}
	if errResp != nil {
		logrus.WithFields(logrus.Fields{
			"error":  errResp.Error,
			"reason": errResp.Reason,
		}).Error("failed to setup user")
		return
	}

	// setup vhosts and permissions
	type rabbitMQSetup struct {
		Vhost     string
		Configure string
		Read      string
		Write     string
	}

	rabbitMQSetups := []rabbitMQSetup{
		{
			Vhost:     rabbitmq.HeartbeatVhost,
			Configure: rabbitmq.HeartbeatConfigurePermissions,
			Read:      rabbitmq.HeartbeatMinionReadPermissions,
			Write:     rabbitmq.HeartbeatMinionWritePermissions,
		},
		{
			Vhost:     rabbitmq.TaskRequestVhost,
			Configure: rabbitmq.TaskRequestConfigurePermissions,
			Read:      rabbitmq.TaskRequestMinionReadPermissions,
			Write:     rabbitmq.TaskRequestMinionWritePermissions,
		},
		{
			Vhost:     rabbitmq.TaskResponseVhost,
			Configure: rabbitmq.TaskResponseConfigurePermissions,
			Read:      rabbitmq.TaskResponseMinionReadPermissions,
			Write:     rabbitmq.TaskResponseMinionWritePermissions,
		},
		{
			Vhost:     rabbitmq.WorkerStatusVhost,
			Configure: rabbitmq.WorkerStatusConfigurePermissions,
			Read:      rabbitmq.WorkerStatusMinionReadPermissions,
			Write:     rabbitmq.WorkerStatusMinionWritePermissions,
		},
		{
			Vhost:     rabbitmq.WorkerEnrollVhost,
			Configure: rabbitmq.WorkerEnrollConfigurePermissions,
			Read:      rabbitmq.WorkerEnrollReadPermissions,
			Write:     rabbitmq.WorkerEnrollWritePermissions,
		},
	}

	for _, setup := range rabbitMQSetups {
		errResp, err = managementClient.Vhosts.Put(
			setup.Vhost,
			"",
			nil,
			vhosts.QueueTypeClassic,
		)
		if err != nil {
			logrus.WithError(err).Error("failed to setup vhost")
			return
		}
		if errResp != nil {
			logrus.WithFields(logrus.Fields{
				"error":  errResp.Error,
				"reason": errResp.Reason,
				"vhost":  setup.Vhost,
			}).Error("failed to setup vhost")
			return
		}

		errResp, err = managementClient.Permissions.Put(
			config.RabbitMQ.Minion.User,
			setup.Vhost,
			setup.Configure,
			setup.Read,
			setup.Write,
		)
		if err != nil {
			logrus.WithError(err).
				WithFields(logrus.Fields{
					"user":      config.RabbitMQ.Minion.User,
					"vhost":     setup.Vhost,
					"configure": setup.Configure,
					"read":      setup.Read,
					"write":     setup.Write,
				}).Error("failed to setup permissions")
			return
		}
		if errResp != nil {
			logrus.WithFields(logrus.Fields{
				"error":     errResp.Error,
				"reason":    errResp.Reason,
				"user":      config.RabbitMQ.Minion.User,
				"vhost":     setup.Vhost,
				"configure": setup.Configure,
				"read":      setup.Read,
				"write":     setup.Write,
			}).Error("failed to setup permissions")
			return
		}
	}

	rabbitmq.Serve(
		ctx,
		taskRequestChan,
		taskResponseChan,
		workerStatusChan,
		redisClient,
		entClient,
	)
}

// serverRun runs the server
func run(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()

	entClient, err := data.NewClient(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create ent client")
	}

	taskRequestChan := make(chan *structs.TaskRequest)
	taskResponseChan := make(chan *structs.TaskResponse)
	workerStatusChan := make(chan *structs.WorkerStatus)
	defer close(taskRequestChan)
	defer close(taskResponseChan)
	defer close(workerStatusChan)

	redisClient := cache.NewRedisClient()

	engineClient := engine.NewEngine(ctx, entClient, redisClient, taskRequestChan, taskResponseChan, workerStatusChan)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go startWebServer(wg, entClient, redisClient, engineClient, taskRequestChan, taskResponseChan, workerStatusChan)
	go startRabbitMQServer(wg, cmd.Context(), taskRequestChan, taskResponseChan, workerStatusChan, redisClient, entClient)

	wg.Wait()

	ctx.Done()

	logrus.Info("Server stopped")
}
