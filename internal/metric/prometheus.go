package metric

import "github.com/prometheus/client_golang/prometheus"

type Service struct {
	httpRequestHistogram *prometheus.HistogramVec
}

func NewPrometheusService() (*Service, error) {
	http := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "http",
			Name:      "request_duration_seconds",
			Help:      "Latency of the HTTP requests",
			Buckets:   prometheus.DefBuckets,
		}, []string{"handler", "method", "code"})

	s := &Service{
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
