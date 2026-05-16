package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const serviceName = "todo-api"

type Metrics struct {
	httpRequestCounter         *prometheus.CounterVec
	httpRequestDurationSeconds *prometheus.HistogramVec
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
		startTime := time.Now()
		next.ServeHTTP(statusRecorder, r)
		if r.URL.Path != "/metrics" {
			labels := prometheus.Labels{
				"service":     serviceName,
				"path":        r.URL.Path,
				"method":      r.Method,
				"status_code": strconv.Itoa(statusRecorder.statusCode),
			}
			duration := time.Since(startTime).Seconds()
			m.httpRequestCounter.With(labels).Inc()
			m.httpRequestDurationSeconds.With(labels).Observe(duration)
		}

	})
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		httpRequestCounter:         httpRequestCounterMetric(),
		httpRequestDurationSeconds: httpRequestDurationSecondsMetric(),
	}
	reg.MustRegister(m.httpRequestCounter)
	reg.MustRegister(m.httpRequestDurationSeconds)
	return m

}

func httpRequestCounterMetric() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Count of http requests",
		},
		[]string{"service", "path", "method", "status_code"},
	)
}

func httpRequestDurationSecondsMetric() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of http request in seconds",
		},
		[]string{"service", "path", "method", "status_code"},
	)
}
