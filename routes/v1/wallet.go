package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddWalletRoutes(r chi.Router) chi.Router {
	r.Route("/wallet", func(r chi.Router) {
		r.Get("/{id}", handler.GetWalletHandler)
		r.Put("/wallet", handler.UpdateWalletHandler)
	})
	return r
}
