package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func setupRouter(port string) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handleErr)

	// Applying the API key middleware only to the /webhook route
	v1Router.With(apiKeyMiddleware).Post("/webhook", handleWebhookEvent)

	router.Mount("/v1", v1Router)

	return router
}
