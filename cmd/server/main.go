package main

import (
	"log"
	"net/http"

	"github.com/shuvo-paul/email-microservice/internal/handlers"
)

func main() {
	http.HandleFunc("GET /health", handlers.HealthHandler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
