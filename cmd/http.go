package cmd

import (
	myhttp "gotravel/cmd/http"
	"gotravel/pkg/middlewares"
	"gotravel/pkg/stdio"
	"gotravel/pkg/validator"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func HTTP(io stdio.IO) {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	io.Echo("# Inicializando HTTP server na porta %s", port)

	validator := validator.New()

	router := chi.NewRouter()

	// midlewares
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Timeout(5 * time.Second))
	router.Use(middlewares.WithIO(io))
	router.Use(middlewares.WithValidator(validator))

	// routes
	myhttp.NewRouter(router)

	// serve
	if err := http.ListenAndServe(":"+port, router); err != nil {
		io.Echo("# Um erro aconteceu: %s\n", err)
	}

	io.Echo("# Servidor encerrado")
}
