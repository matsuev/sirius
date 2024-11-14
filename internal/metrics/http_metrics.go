package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
