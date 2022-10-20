package handlers

import (
	"github.com/go-chi/chi"
)

func RegisterRoutes(mux *chi.Mux) *chi.Mux {
	mux = registerHealthRoutes(mux)
	mux = registerUserRoutes(mux)
	mux = registerSecretRoutes(mux)
	return mux
}

func registerHealthRoutes(mux *chi.Mux) *chi.Mux {
	mux.Get("/", DefaultHandler)
	mux.Get("/health", HealthHandler)
	mux.Get("/health/full", FullHealthHandler)
	return mux
}

func registerUserRoutes(mux *chi.Mux) *chi.Mux {
	mux.Get("/users", GetUsersHandler)
	mux.Post("/users", CreateUserHandler)
	mux.Get("/users/{user_id}", GetUserHandler)
	return mux
}

func registerSecretRoutes(mux *chi.Mux) *chi.Mux {
	mux.Get("/secrets", GetAllSecretsHandler)
	mux.Post("/secrets", CreateSecretHandler)
	mux.Get("/secrets/{secret_id}", GetSecretHandler)
	mux.Post("/secrets/{secret_id}/guess", GuessSecretHandler)
	return mux
}
