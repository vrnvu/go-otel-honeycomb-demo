package handlers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func NewServerHandler(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServerHandler(w, r, path)
	})
}

func ServerHandler(w http.ResponseWriter, r *http.Request, path string) {
	if path != "foo" && path != "bar" && path != "baz" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Not Found: %s", path)))
		return
	}

	providerBaseURL := os.Getenv("PROVIDER_BASE_URL")
	if providerBaseURL == "" {
		providerBaseURL = "http://localhost:8081"
	}

	resp, err := otelhttp.Get(r.Context(), providerBaseURL+"/"+path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	var statusCode int
	var statusMessage string
	switch resp.StatusCode {
	case http.StatusOK:
		statusCode = http.StatusOK
		statusMessage = "OK"
	case http.StatusBadRequest:
		statusCode = http.StatusBadRequest
		statusMessage = "Bad Request"
	case http.StatusServiceUnavailable:
		statusCode = http.StatusServiceUnavailable
		statusMessage = "Service Unavailable"
	default:
		statusCode = http.StatusInternalServerError
		statusMessage = "Internal Server Error"
	}

	_, span := otel.Tracer("ServerHandlerBody").Start(r.Context(), "ServerHandlerBody")

	randomSleep := rand.IntN(100)
	sleepDuration := time.Duration(randomSleep) * time.Millisecond

	span.AddEvent("starting body processing")
	time.Sleep(sleepDuration)
	span.AddEvent("body processing complete")

	span.End()

	dbBaseURL := os.Getenv("DB_BASE_URL")
	if dbBaseURL == "" {
		dbBaseURL = "http://localhost:8082"
	}

	resp, err = otelhttp.Get(r.Context(), dbBaseURL+"/")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(statusCode)
	w.Write([]byte(statusMessage))
}
