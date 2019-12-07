package routesv1

import (
	"github.com/go-chi/chi"
	"grabathon/handler"
)

func AddUserRoutes(r chi.Router) chi.Router {
	r.Route("/user", func(r chi.Router) {
		r.Get("/users", handler.GetAllUsersHandler)
		r.Get("/{id}", handler.GetUserHandler)
		//r.Put("/", handler.UpdateUserHandler)
	})

	r.Post("/user/create", handler.CreateUserHandler)
	return r
}
