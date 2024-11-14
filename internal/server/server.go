package server

import (
	"fmt"
	"log/slog"

	"sirius-metrics/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// Listen ...
func Listen(addr string) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", handlers.HomeHandler())
	app.Get("/help", handlers.HelpHandler())

	slog.Info(fmt.Sprintf("api server available on http://%s/", addr))

	return app.Listen(addr)
}
