package routes

import (
	"github.com/go-chi/chi"
	"grabathon/middleware"
	routesv1 "grabathon/routes/v1"
)

/**
 * Created by Sai Ravi Teja K on 22, Jul 2019
 */

// Add new routes to index them
func ConfigureRoutes(router *chi.Mux) chi.Router {
	router.Route("/", func(r chi.Router) {
		routesv1.AddHealthCheckRoutes(r)
	})
	router.Group(func(router chi.Router){
		router.Use(middleware.Authenticate)
		router.Route("/api", func(r chi.Router) {
			routesv1.AddUserRoutes(r)
			routesv1.AddWalletRoutes(r)
			routesv1.AddVouchRoutes(r)
		})
	})
	return router
}