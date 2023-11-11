package server

import (
	"log"
	"os"
	"os/signal"
	"project/go-fiber-boilerplate/config"
	"project/go-fiber-boilerplate/infrastructure/database"
	"project/go-fiber-boilerplate/infrastructure/http/handler/auth"
	"project/go-fiber-boilerplate/infrastructure/http/handler/user"
	"project/go-fiber-boilerplate/utils/constants"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App  *fiber.App
	Conf *config.AppConfig
}

func NewServer() *Server {
	app := fiber.New()
	conf := config.NewAppConfig()
	return &Server{
		App:  app,
		Conf: conf,
	}

}

func (s *Server) Run() {
	conf := s.Conf
	app := s.App

	db := database.NewPostgres(conf)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Authorization, Content-Length",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.ConfigDefault))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(
			constants.NewSuccess("welcome to go fiber boilerplate", nil),
		)
	})

	// routes
	auth.NewAuthRoutes(app, db)
	user.NewUserRoutes(app, db)

	s.GracefulShutdown(conf.Fiber.Port)
}

func (s *Server) GracefulShutdown(port string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.App.Listen(":" + port); err != nil {
			log.Fatalf("error when listening to :%s, %s", port, err)
		}
	}()

	log.Printf("server is running on :%s", port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := s.App.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}
