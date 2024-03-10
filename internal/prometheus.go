package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PrometheusInstance struct {
	requestsTotal    *prometheus.CounterVec
	requestsInFlight *prometheus.GaugeVec
	url              string
}

func StartPrometheus() *PrometheusInstance {
	constLabels := prometheus.Labels{
		"app": "fiber-prometheus-demo",
	}

	requests_total_counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "requests_total"),
			Help:        "The total number of HTTP requests made",
			ConstLabels: constLabels,
		},
		[]string{"status_code", "method", "path"},
	)

	requests_in_flight_gauge := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "requests_in_flight_total"),
			Help:        "The total number of HTTP requests in flight",
			ConstLabels: constLabels,
		},
		[]string{"method"},
	)

	return &PrometheusInstance{
		requestsTotal:    requests_total_counter,
		requestsInFlight: requests_in_flight_gauge,
		url:              "/metrics",
	}
}

func (p *PrometheusInstance) Middleware(c *fiber.Ctx) error {
	method := c.Route().Method
	// path := c.Route().Path
	p.requestsInFlight.WithLabelValues(method).Inc()
	defer func() {
		p.requestsInFlight.WithLabelValues(method).Dec()
	}()
	return c.Next()
}
