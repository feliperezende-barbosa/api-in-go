package metric

import "time"

// HTTP Application
type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

// NewHTTP creates a new HTTP app
func NewHTTP(handler string, method string) *HTTP {
	return &HTTP{
		Handler: handler,
		Method:  method,
	}
}

// Started starts monitoring the app
func (h *HTTP) Started() {
	h.StartedAt = time.Now()
}

// Started finishes monitoring the app
func (h *HTTP) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}

type UseCase interface {
	SaveHTTP(h *HTTP)
}
