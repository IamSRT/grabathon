package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddUserRoutes(r chi.Router) chi.Router {
	r.Get("/user", handler.GetAllUsersHandler)
	r.Post("/user", handler.CreateUserHandler)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handler.GetUserHandler)
	})
	return r
}
