package middleware

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const serviceName = "todo-api"

type Metrics struct {
	httpRequestCounter *prometheus.CounterVec
}

type StatusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *StatusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (m *Metrics) MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusRecorder := &StatusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(statusRecorder, r)
		if r.URL.Path != "/metrics" {
			m.httpRequestCounter.WithLabelValues(serviceName, r.URL.Path, r.Method, strconv.Itoa(statusRecorder.statusCode)).Inc()
		}

	})
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := httpRequestCounterMetric()
	reg.MustRegister(m.httpRequestCounter)
	return m

}

func httpRequestCounterMetric() *Metrics {
	return &Metrics{
		httpRequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Count of http requests",
			},
			[]string{"service", "path", "method", "status_code"},
		),
	}
}
