package rest

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func withLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("handling request",
			"uri", r.RequestURI,
			"method", r.Method,
		)

		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		h.ServeHTTP(&loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}, r)

		duration := time.Since(start)

		slog.Debug("handled request",
			"uri", r.RequestURI,
			"method", r.Method,
			"status", responseData.status,
			"duration", duration,
			slog.Group("size",
				"bytes", responseData.size,
				"readable", humanize.Bytes(uint64(responseData.size)),
			),
		)
	})
}
