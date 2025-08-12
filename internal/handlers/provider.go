package handlers

import (
	"math/rand/v2"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewProviderHandler(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch path {
		case "/foo":
			ProviderHandler(w, r, "foo")
		case "/bar":
			ProviderHandler(w, r, "bar")
		case "/baz":
			ProviderHandler(w, r, "baz")
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func ProviderHandler(w http.ResponseWriter, r *http.Request, path string) {
	span := trace.SpanFromContext(r.Context())
	span.SetAttributes(attribute.String("metrics.provider.path", path))

	randomSleep := rand.IntN(100)
	time.Sleep(time.Duration(randomSleep) * time.Millisecond)

	if randomSleep < 80 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(path))
	} else if randomSleep < 95 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
