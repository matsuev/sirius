package metrics

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Listen(addr string) error {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	slog.Info(fmt.Sprintf("metrics server available on http://%s/metrics", addr))

	return http.ListenAndServe(addr, mux)
}
