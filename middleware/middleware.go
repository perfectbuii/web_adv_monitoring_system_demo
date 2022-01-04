package middleware

import (
	"fmt"
	"net/http"
	"time"
	"web_adv/metrics"
	"web_adv/utils"
)

const (
	MetricPath = "/metrics"
)

func WithMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &utils.ResponseWriter{
			ResponseWriter: w,
		}
		if r.URL.Path != MetricPath {
			start := time.Now()
			defer func() {
				metrics.OnRequestFinish(r, rw.Status)
				metrics.OnCalcFinish(fmt.Sprintf("%v.%v", r.Method, r.URL.Path), time.Since(start))
			}()
		}
		next.ServeHTTP(rw, r)
	})
}
