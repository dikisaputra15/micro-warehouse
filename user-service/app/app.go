package app

import (
	"context"
	"fmt"
	"micro-warehouse/user-service/configs"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() { 
	cfg := configs.NewConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Errorf("Error: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		},
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] $ip ${status} - ${latency} ${method} ${path}\n",
	}))

	container := BuildContainer()
	SetupRoutes(app, container)

	port := cfg.App.AppPort
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			log.Fatalf("Server port not specified")
		}
	}
	log.Infof("Starting server on port %s", port)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("Error Starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}
	log.Info("Server shutdown complete")
}
