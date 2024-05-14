package rest

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

type requestContextKey struct{}
type RequestContext map[string]string

func withRequestContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, requestContextKey{}, RequestContext{
			"id": uuid.New().String(),
		})
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

type requestContextHandler struct {
	slog.Handler
}

func (h requestContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestCtx, ok := ctx.Value(requestContextKey{}).(RequestContext); ok {
		attrs := make([]slog.Attr, 0, len(requestCtx))
		for k, v := range requestCtx {
			attrs = append(attrs, slog.String(k, v))
		}
		r.Add("request", slog.GroupValue(attrs...))
	}

	return h.Handler.Handle(ctx, r)
}

func NewRequestContextSlogHandler(h slog.Handler) slog.Handler {
	return requestContextHandler{h}
}

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
		slog.DebugContext(r.Context(), "handling request",
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

		slog.DebugContext(r.Context(), "handled request",
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
