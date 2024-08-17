package rest

import "github.com/go-chi/chi"

func (s *Server) routes() {
	// orders group
	s.Router.Route("/orders", func(r chi.Router) {
		r.Get("/", s.ListOrders)
		r.Post("/", s.CreateOrder)
	})

	// items group
	s.Router.Route("/items", func(r chi.Router) {
		r.Get("/", s.ListItems)
	})

}
