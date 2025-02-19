package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"method", "endpoint"},
	)

	OrdersProcessGoroutinesTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "orders_process_goroutines_total",
			Help: "Total number of goroutines process orders",
		},
	)

	UnprocessedItensTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "unprocessed_itens_total",
			Help: "Total number of itens not processed",
		},
	)

	ProcessedItensTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "processed_itens_total",
			Help: "Total number of itens processed",
		},
	)

	ItensTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "itens_total",
			Help: "Total number of itens",
		},
	)

	OrdersTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "orders_total",
			Help: "Total number of orders",
		},
	)
	NotFoundOrders = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "orders_not_found_total",
			Help: "Total number of not found orders",
		},
	)

	TotalRequestDC = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "request_dc_api_total",
			Help: "Total number of request into distribution center api",
		},
	)
)

func init() {
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(HttpRequestDuration)
	prometheus.MustRegister(OrdersProcessGoroutinesTotal)
	prometheus.MustRegister(UnprocessedItensTotal)
	prometheus.MustRegister(ProcessedItensTotal)
	prometheus.MustRegister(NotFoundOrders)
	prometheus.MustRegister(OrdersTotal)
	prometheus.MustRegister(ItensTotal)
	prometheus.MustRegister(TotalRequestDC)
}

// func RecordMetrics(method string, endpoint string, duration float64) {
// 	httpRequestsTotal.WithLabelValues(method, endpoint).Inc()
// 	httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)
// }
