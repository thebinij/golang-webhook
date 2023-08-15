package main

import (
	"net/http"
	"os"
)

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		validAPIKey := os.Getenv("VALID_API_KEY") // Use environment variable for the valid API key

		if apiKey != validAPIKey {
			responseWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
