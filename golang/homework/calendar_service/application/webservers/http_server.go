package webservers

import (
	"calendar_service/application/configs"
	"calendar_service/application/services/api/rest"
	"calendar_service/frameworks/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "code"},
	)

	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_latency_seconds",
			Help:    "Latency of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"method", "code"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestLatency)
	prometheus.MustRegister(errorCount)
}

func MetricServer() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	// Start the HTTP server using the custom ServeMux
	logger.Info("Starting metric server on :9001")
	if err := http.ListenAndServe(":9001", mux); err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}

func Server(c *configs.ServerConfig) {

	rest.Config(&c.Storage)
	if err := http.ListenAndServe(c.HTTPAddress, nil); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func RouterInit(c configs.ServerConfig) {

	http.HandleFunc("/create_event", applyMiddleware(rest.CreateEvent, metricsMiddleware, logMiddleware))
	http.HandleFunc("/update_event", applyMiddleware(rest.UpdateEvent, metricsMiddleware, logMiddleware))
	http.HandleFunc("/delete_event", applyMiddleware(rest.DeleteEvent, metricsMiddleware, logMiddleware))
	http.HandleFunc("/events_for_day", applyMiddleware(rest.EventsForDay, metricsMiddleware, logMiddleware))
	http.HandleFunc("/events_for_week", applyMiddleware(rest.EventsForWeek, metricsMiddleware, logMiddleware))
	http.HandleFunc("/events_for_month", applyMiddleware(rest.EventsForMonth, metricsMiddleware, logMiddleware))

}

func applyMiddleware(next http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		next = m(next)
	}
	return next
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next(w, r)

		logger.Info(fmt.Sprintf("%s/%s/%s/%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start)))
	}

}

func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		method := r.Method
		statusCode := rw.statusCode

		requestCount.WithLabelValues(method, http.StatusText(statusCode)).Inc()
		requestLatency.WithLabelValues(method).Observe(duration)

		if statusCode >= 400 {
			logger.Error("Error ", statusCode)
			errorCount.WithLabelValues(method, http.StatusText(statusCode)).Inc()
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
