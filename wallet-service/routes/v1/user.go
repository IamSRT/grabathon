package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddUserRoutes(r chi.Router) chi.Router {
	r.Route("/user", func(r chi.Router) {
		r.Get("/users", handler.GetAllUsersHandler)
		r.Get("/{id}", handler.GetUserHandler)
	})

	r.Put("/user/update", handler.UpdateUserHandler)
	r.Post("/user/create", handler.CreateUserHandler)
	r.Post("/user/vouch", handler.Vouch)
	r.Post("/user/enable-auto_pay", handler.EnableAutoPay)

	r.Post("/is-auto-pay-enabled", handler.IsAutoPayEnabled)
	r.Post("/is-vouch-valid", handler.IsVouchValid)
	r.Get("/user/{id}/pending-payments", handler.GetPendingPaymentsHandler)
	return r
}
