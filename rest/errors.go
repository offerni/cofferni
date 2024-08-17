package rest

import "errors"

var (
	ErrHttpServerIsRequired  error = errors.New("HttpServerIsRequired is required")
	ErrMenuServiceIsRequired error = errors.New("MenuService is required")
	ErrRouterIsRequired      error = errors.New("Router is required")
)

type ErrorResponse struct {
	Error string `json:"error"`
}
