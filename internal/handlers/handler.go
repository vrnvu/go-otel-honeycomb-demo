package handlers

import (
	"net/http"
)

func NewHelloHandler() http.Handler {
	return http.HandlerFunc(HelloHandler)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
