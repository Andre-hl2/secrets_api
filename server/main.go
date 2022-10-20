package main

import (
	"fmt"
	"log"
	"net/http"
	"secret_api/api/handlers"
	"secret_api/config"
	"secret_api/storage"

	"github.com/go-chi/chi"
)

func main() {
	// Get config object
	cfg, err := config.BuildConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Register routes
	mux := handlers.RegisterRoutes(chi.NewRouter())

	// Wrap the routes with the store object
	var contextMux http.Handler

	switch cfg.DataSourceType {
	case config.MemorySource:
		st := storage.NewMemoryStore()

		log.Print("Running service with memory store")
		contextMux = handlers.AddStore(mux, &st)
	case config.PostgresSource:
		st, err := storage.NewPostgresStore(cfg)
		if err != nil {
			err_txt := fmt.Sprintf("Failed to initialize Postgres store: %s", err.Error())
			log.Fatalf(err_txt)
		}

		log.Print("Running service with postgres store")
		contextMux = handlers.AddStore(mux, &st)
	}

	// Listen and serve
	port := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Print(fmt.Sprintf("Starting web server on port %s...", cfg.ServerPort))
	http.ListenAndServe(port, contextMux)
}
