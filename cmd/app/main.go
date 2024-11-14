package main

import (
	"fmt"
	"log/slog"
	"sirius-metrics/internal/metrics"
	"sirius-metrics/internal/server"
)

func main() {
	slog.Info("app started")

	go func() {
		if err := metrics.Listen("127.0.0.1:8082"); err != nil {
			slog.Error("metrics server", slog.Any("err", err))
		}
	}()

	go func() {
		if err := server.Listen("127.0.0.1:8080"); err != nil {
			slog.Error("api server", slog.Any("err", err))
		}
	}()

	var s string
	fmt.Scanln(&s)
}
