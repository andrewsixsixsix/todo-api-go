package middleware

import (
	"log/slog"
	"net/http"
	"time"
	"todo-api/internal/logger"

	"github.com/go-chi/chi/v5/middleware"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqStartTime := time.Now()
		wr := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			duration := time.Since(reqStartTime).String()

			req := slog.Group("request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("headers", r.Header),
			)
			res := slog.Group("response",
				slog.Int("status", wr.Status()),
				slog.Int("bytes", wr.BytesWritten()),
				slog.Any("headers", wr.Header()),
			)

			logger.Logger().Info("processed request", slog.String("duration", duration), req, res)
		}()

		next.ServeHTTP(w, r)
	})
}
