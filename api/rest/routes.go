package rest

import "github.com/go-chi/chi"

func (s *Server) routes() {
	s.Router.Route("/api", func(r chi.Router) {

		// orders group
		r.Route("/orders", func(r chi.Router) {
			r.Get("/", s.ListOrders)
			r.Post("/", s.CreateOrder)
			r.Patch("/", s.UpdateOrder)
		})

		// items group
		r.Route("/items", func(r chi.Router) {
			r.Get("/", s.ListItems)
		})
	})
}
