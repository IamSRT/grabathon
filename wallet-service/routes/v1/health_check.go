package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

func AddHealthCheckRoutes(router chi.Router) chi.Router {
	router.Get("/health", handler.HealthCheckHandler)
	return router
}