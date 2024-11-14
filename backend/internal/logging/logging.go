package logging

import (
	"log/slog"
	"os"
)

// Init ...
func Init() {
	serviceName := os.Getenv("SERVICE")
	if serviceName == "" {
		serviceName = "undefined"
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})).With(slog.String("service", serviceName)))
}
