package routes

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
	"grabathon/middleware"
	routesv1 "grabathon/routes/v1"
)

/**
 * Created by Sai Ravi Teja K on 22, Jul 2019
 */

func AddHealthCheckRoutes(router chi.Router) chi.Router {
	router.Get("/health", handler.HealthCheckHandler)
	return router
}

// Add new routes to index them
func ConfigureRoutes(router *chi.Mux) chi.Router {
	router.Route("/", func(r chi.Router) {
		AddHealthCheckRoutes(r)
	})
	router.Group(func(router chi.Router){
		router.Use(middleware.Authenticate)
		router.Route("/api/v1", func(r chi.Router) {
			routesv1.AddStatsRoutes(r)
		})
	})
	return router
}