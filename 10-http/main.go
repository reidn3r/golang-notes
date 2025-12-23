package main

import (
	"learning/net/handlers"
	"net/http"
)

func main() {
	http.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	http.ListenAndServe(":8080", nil)
}
