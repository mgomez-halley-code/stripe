package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newResponseWriter creates a new responseWriter
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, // default status code
	}
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging logs HTTP requests with method, path, status code, and duration
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap the response writer to capture status code
		wrappedWriter := newResponseWriter(w)

		// Call the next handler
		next.ServeHTTP(wrappedWriter, r)

		// Log the request details
		duration := time.Since(start)
		log.Printf(
			"%s %s - Status: %d - Duration: %v",
			r.Method,
			r.URL.Path,
			wrappedWriter.statusCode,
			duration,
		)
	})
}
