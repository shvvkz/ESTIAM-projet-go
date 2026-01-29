package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// responseWriter
//
// A custom HTTP response writer that captures the status code of the response.
//
// Fields:
//   - ResponseWriter: The original http.ResponseWriter.
//   - status: The HTTP status code of the response.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader(code int)
//
// Captures the status code and calls the original WriteHeader method.
//
// Parameters:
//   - code: The HTTP status code to write.
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger(next http.Handler) -> http.Handler
//
// Middleware that logs HTTP requests, including method, path, status code, and duration.
//
// Parameters:
//   - next: The next HTTP handler in the chain.
//
// Returns:
//   - An http.Handler that wraps the next handler with logging functionality.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		fmt.Printf(
			"[HTTP] %s %s -> %d (%s)\n",
			r.Method,
			r.URL.Path,
			rw.status,
			duration,
		)
	})
}
