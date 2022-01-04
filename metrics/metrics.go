package metrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "web_adv",
			Subsystem:   "test",
			Name:        "request_total",
			Help:        "",
			ConstLabels: nil,
		},
		[]string{"method", "path", "status"},
	)
	CalcTimeHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "web_adv",
			Subsystem:   "test",
			Name:        "request_time",
			Help:        "",
			ConstLabels: nil,
			Buckets: []float64{
				0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9,
				1, 2, 3, 4, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60,
				120, 240,
			},
		},
		[]string{"process_func_time"},
	)
)

func OnRequestFinish(req *http.Request, status int) {
	RequestCounter.WithLabelValues(req.Method, req.URL.Path, fmt.Sprint(status)).Inc()
}

func OnCalcFinish(funcName string, duration time.Duration) {
	CalcTimeHistogram.WithLabelValues(funcName).Observe(duration.Seconds())
}
