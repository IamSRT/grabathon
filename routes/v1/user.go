package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddUserRoutes(r chi.Router) chi.Router {
	r.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetAllUsersHandler)
		r.Get("/{id}", handler.GetUserHandler)
		r.Post("/user", handler.CreateUserHandler)
		r.Put("/user", handler.UpdateUserHandler)
	})
	return r
}
