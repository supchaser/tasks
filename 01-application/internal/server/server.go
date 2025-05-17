package server

import (
	"net/http"
	"time"

	"CLOUDRU/internal/config"
	"CLOUDRU/internal/handler"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var ErrServerClosed = http.ErrServerClosed

func New(cfg *config.Config, log *logrus.Logger) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.EchoHandler(log, cfg.Author)).Methods(http.MethodGet)
	router.HandleFunc("/healthz", handler.HealthHandler(log)).Methods(http.MethodGet)

	router.Use(loggingMiddleware(log))

	return &http.Server{
		Addr:         cfg.Addr(),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func loggingMiddleware(log *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.WithFields(logrus.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"remote": r.RemoteAddr,
				"took":   time.Since(start).String(),
			}).Info("handled request")
		})
	}
}
