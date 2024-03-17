package internal

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// this file was based on https://github.com/ansrivas/fiberprometheus

type PrometheusInstance struct {
	requestsTotal     *prometheus.CounterVec
	requestsInFlight  *prometheus.GaugeVec
	requestTimes      *prometheus.HistogramVec
	APITriggersMetric *prometheus.CounterVec
	url               string
}

var Prometheus *PrometheusInstance

func StartPrometheus() *PrometheusInstance {
	constLabels := prometheus.Labels{
		"app": "fiber-prometheus-demo",
	}

	// this metric will store all requests sent to the server. Use this to get the rate of requests per minute or second
	requests_total_counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "requests_total"),
			Help:        "The total number of HTTP requests made",
			ConstLabels: constLabels,
		},
		[]string{"status_code", "method", "path"},
	)

	// this metric will count how many requests are being processed at a specific instant
	requests_in_flight_gauge := promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "requests_in_flight_total"),
			Help:        "The total number of HTTP requests in flight",
			ConstLabels: constLabels,
		},
		[]string{"method"},
	)

	// this metric will count how long each request takes to process
	request_timing_histogram := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "request_duration_seconds"),
			Help:        "The duration of HTTP requests in seconds",
			ConstLabels: constLabels,
			Buckets:     []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"endpoint"},
	)

	// this metric will count how many requests were sent to the trigger XXX messag endpoints in api/metricstesting_handlers.go
	// tbh, this one isn't that useful like this, but I could prolly use something like this to measure how much a specific function is called.
	// maybe if it's a heavy function that's constantly being called, we could evaluate taking it to AWS Lambda for example?
	some_metric_counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName("my_go_project", "api", "some_metric"),
			Help:        "Some metric",
			ConstLabels: constLabels,
		},
		[]string{"triggered_handler"},
	)

	Prometheus = &PrometheusInstance{
		requestsTotal:     requests_total_counter,
		requestsInFlight:  requests_in_flight_gauge,
		APITriggersMetric: some_metric_counter,
		requestTimes:      request_timing_histogram,
		url:               "/metrics",
	}

	return Prometheus
}

// this is the middleware that will run request-based metrics
func (p *PrometheusInstance) Middleware(c *fiber.Ctx) error {
	startTime := time.Now()
	method := c.Route().Method
	path := c.Route().Path

	p.requestsInFlight.WithLabelValues(method).Inc()
	defer func() {
		p.requestsInFlight.WithLabelValues(method).Dec()
	}()

	var status int

	err := c.Next()
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			// Get correct error code from fiber.Error type
			status = e.Code
		}
	} else {
		status = c.Response().StatusCode()
	}

	endTime := float64(time.Since(startTime).Microseconds()) / 1e6

	p.requestTimes.WithLabelValues(path).Observe(endTime)

	statusCodeString := strconv.Itoa(status)
	p.requestsTotal.WithLabelValues(statusCodeString, method, path).Inc()

	return err
}
