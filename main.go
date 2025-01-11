package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api", apiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request on %s", r.URL.Path)

	// Print the X-Request-From header
	requestFrom := r.Header.Get("X-Request-From")
	if requestFrom == "" {
		requestFrom = "Unknown"
	}
	log.Printf("X-Request-From: %s", requestFrom)

	// Respond to the request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from %s\n", os.Getenv("APP_NAME"))
	fmt.Fprintf(w, "Received from: %s\n", requestFrom)
}
