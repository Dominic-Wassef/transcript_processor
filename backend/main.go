package main

import (
	"backend/handlers"
	"backend/middleware"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/api/transcript", handlers.TranscriptHandler)

	// middleware
	wrappedMux := middleware.CorsMiddleware(mux)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
