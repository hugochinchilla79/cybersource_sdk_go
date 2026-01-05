package cybersource

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
	Headers    http.Header
	RequestID  string
}

func (e *HTTPError) Error() string {
	if len(e.Body) > 0 {
		return fmt.Sprintf("cybersource http error %d (%s): %s", e.StatusCode, e.Status, string(e.Body))
	}
	return fmt.Sprintf("cybersource http error %d (%s)", e.StatusCode, e.Status)
}
