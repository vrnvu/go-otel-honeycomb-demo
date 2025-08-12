package handlers

import (
	"math/rand/v2"
	"net/http"
	"time"
)

func NewDBHandler() http.Handler {
	return http.HandlerFunc(DBHandler)
}

func DBHandler(w http.ResponseWriter, r *http.Request) {
	randomSleep := rand.IntN(100)
	time.Sleep(time.Duration(randomSleep) * time.Millisecond)

	if randomSleep < 80 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
