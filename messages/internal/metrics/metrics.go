package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"time"
)

type Metrics interface {
	ObserveRequest(name string, time time.Duration, status int)
}

func NewMetrics() Metrics {
	http.NewServeMux().Handle("/metrics", promhttp.Handler())
	return &metrics{
		requestsObserve: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "requests",
		}, []string{"name", "status"}),
	}
}

type metrics struct {
	requestsObserve *prometheus.HistogramVec
}

func (m metrics) ObserveRequest(name string, time time.Duration, status int) {
	m.requestsObserve.With(prometheus.Labels{"name": name, "status": strconv.Itoa(status)}).Observe(time.Seconds())
}
