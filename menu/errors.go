package menu

import "errors"

var (
	ErrItemRepositoryIsRequired  error = errors.New("ItemRepository is required")
	ErrOrderRepositoryIsRequired error = errors.New("OrderRepository is required")
)
