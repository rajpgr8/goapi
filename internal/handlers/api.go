package handlers

import (
	"fmt"

	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	fmt.Println("Starting Handler...")

	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)

	})
}
