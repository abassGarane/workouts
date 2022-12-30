package main

import (
	"fmt"
	"net/http"
)

var Calls int64

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Calls += 1
		fmt.Println(Calls)
		next.ServeHTTP(w,r)
	})
}
