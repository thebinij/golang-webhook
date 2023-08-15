package main

import (
	"net/http"
	"os"
)

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			responseWithError(w, http.StatusUnauthorized, "Unauthorized: API KEY")
			return
		}
		validAPIKey := os.Getenv("VALID_API_KEY")
		if apiKey != validAPIKey {
			responseWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
