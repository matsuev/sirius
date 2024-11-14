package main

import (
	"log"
	"log/slog"
	"logging-stack/backend/internal/logging"
	"logging-stack/backend/internal/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	slogfiber "github.com/samber/slog-fiber"
)

var requestTotals = promauto.NewCounterVec(prometheus.CounterOpts{
	Namespace: "server",
	Subsystem: "http",
	Name:      "requests_total_with_path",
}, []string{"url"})

// ObserveRequest ...
func ObserveRequest(path string) {
	requestTotals.WithLabelValues(path).Inc()
}

func main() {
	logging.Init()

	go func() {
		if err := metrics.Listen(":8082"); err != nil {
			slog.Error("metrics server", slog.Any("err", err))
		}
	}()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(slogfiber.New(slog.Default()))
	app.Use(recover.New())

	rpc := app.Group("/rpc")

	rpc.Get("/", func(c *fiber.Ctx) error {
		ObserveRequest(c.Path())

		return c.SendString("Hello, World 👋!")
	})

	if err := app.Listen(":80"); err != nil {
		log.Fatalln(err)
	}
}
