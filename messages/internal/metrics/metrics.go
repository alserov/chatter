package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

type Metrics interface {
	ObserveRequest(name string, time time.Duration, status int)
}

func NewMetrics() Metrics {
	return &metrics{}
}

type metrics struct {
	requestsObserve prometheus.HistogramVec
}

func (m metrics) ObserveRequest(name string, time time.Duration, status int) {
	m.requestsObserve.With(prometheus.Labels{"name": name, "status": strconv.Itoa(status)}).Observe(time.Seconds())
}
