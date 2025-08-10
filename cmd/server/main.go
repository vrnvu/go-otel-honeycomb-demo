package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vrnvu/go-otel-honeycomb-demo/internal/handlers"
	"github.com/vrnvu/go-otel-honeycomb-demo/internal/telemetry"
)

func main() {
	otelShutdown, err := telemetry.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("failed to configure OpenTelemetry: %v", err)
	}
	defer otelShutdown()

	router := http.NewServeMux()
	router.Handle("/", telemetry.NewHandler(handlers.NewHelloHandler(), "HelloHandler"))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited")
}
