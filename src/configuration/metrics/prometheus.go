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
		[]string{"method", "endpoint", "status"},
	)
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"method", "endpoint", "status"},
	)

	OrdersProcessGoroutinesTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "orders_process_goroutines_total",
			Help: "Total number of workers process orders",
		},
	)

	UnprocessedItemsTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "unprocessed_items_total",
			Help: "Total number of items not processed",
		},
	)

	ProcessedItemsTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "processed_items_total",
			Help: "Total number of items processed",
		},
	)

	ItemsTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "items_total",
			Help: "Total number of items",
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
	prometheus.MustRegister(UnprocessedItemsTotal)
	prometheus.MustRegister(ProcessedItemsTotal)
	prometheus.MustRegister(NotFoundOrders)
	prometheus.MustRegister(OrdersTotal)
	prometheus.MustRegister(ItemsTotal)
	prometheus.MustRegister(TotalRequestDC)
}
