package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		duration := time.Since(start).Seconds()

		method := c.Request.Method
		endpoint := c.FullPath()
		statusCode := strconv.Itoa(c.Writer.Status())

		HttpRequestsTotal.WithLabelValues(method, endpoint, statusCode).Inc()
		HttpRequestDuration.WithLabelValues(method, endpoint, statusCode).Observe(duration)
	}
}
