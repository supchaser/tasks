package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"CLOUDRU/internal/config"
	"CLOUDRU/internal/logger"
	"CLOUDRU/internal/server"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv := server.New(cfg, log)

	go func() {
		log.Infof("starting server on %s", cfg.Addr())
		if err := srv.ListenAndServe(); err != nil && err != server.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}
	log.Info("server gracefully stopped")
}
