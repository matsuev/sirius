package main

import (
	"log/slog"
	"logging-stack/backend/internal/logging"
	"time"
)

func main() {
	logging.Init()

	for {
		slog.Debug("debug message", slog.String("msg", "test debug logging"))
		time.Sleep(2 * time.Second)

		slog.Info("info message", slog.String("msg", "test info logging"))
		time.Sleep(2 * time.Second)

		slog.Error("error message", slog.String("msg", "test error logging"))
		time.Sleep(2 * time.Second)
	}

}
