package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddVouchRoutes(r chi.Router) chi.Router {
	r.Route("/vouch", func(r chi.Router) {
		r.Get("/{id}", handler.GetVouchHandler)
		r.Get("/{id}/vouchers", handler.GetAllVouchesForVoucheeHandler)
		r.Get("/{id}/vouches", handler.GetAllVouchesForVoucherHandler)
	})

	r.Post("/vouch/create", handler.CreateVouchHandler)
	r.Post("/vouch/{id}/accept", handler.AcceptVouch)
	r.Post("/vouch/{id}/reject", handler.RejectVouch)
	r.Post("/vouch/update", handler.UpdateVouchHandler)
	r.Post("/vouches/create", handler.CreateVouchesHandler)
	r.Delete("/vouch/{id}", handler.DeleteVouchHandler)
	return r
}
