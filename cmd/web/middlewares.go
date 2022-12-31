package main

import (
	"fmt"
	"net/http"
)

var Calls int64

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		Calls += 1
		fmt.Println(Calls)
		next.ServeHTTP(w, r)
	})
}
