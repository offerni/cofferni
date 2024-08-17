package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/offerni/cofferni/menu"
)

type Server struct {
	MenuService *menu.Service
	Router      *chi.Mux
	httpServer  *http.Server
}

type NewServerOpts struct {
	MenuService *menu.Service
	Port        string
	Router      *chi.Mux
}

func NewServer(opts NewServerOpts) (*Server, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	server := &Server{
		Router:      opts.Router,
		MenuService: opts.MenuService,
	}

	server.Router.Use(middleware.Timeout(30 * time.Second))

	server.routes()

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", opts.Port),
		Handler: server.Router,
	}

	return server, nil
}

func (s *Server) Start() error {
	fmt.Printf("Server started on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Gracefully shutting down server...")
	return s.httpServer.Shutdown(ctx)
}

func (opts NewServerOpts) Validate() error {
	if opts.Router == nil {
		return ErrRouterIsRequired
	}

	if opts.MenuService == nil {
		return ErrMenuServiceIsRequired
	}

	return nil
}
