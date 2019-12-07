package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddPaymentRoutes(r chi.Router) chi.Router {
	r.Route("/payment", func(r chi.Router) {
		r.Get("/{id}", handler.GetPaymentHandler)
	})

	r.Post("/payment/create", handler.CreatePaymentHandler)
	r.Delete("/{id}", handler.DeletePaymentHandler)
	return r
}
