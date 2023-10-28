package api

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", SearchUsers)
				r.Post("/", CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", GetUser)
					r.Patch("/", UpdateUser)
					r.Delete("/", DeleteUser)
				})
			})
		})
	})

	return r
}
