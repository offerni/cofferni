package cofferni

import (
	"context"
)

type Order struct {
	ID          string
	ItemID      string
	Observation *string
	Quantity    uint
}

type OrderCreateOpts struct {
	ItemID      string
	Observation *string
	Quantity    uint
}

type OrderFindAllOpts struct {
}

type OrderList struct {
	Data []Order
	// pagintion later maybe
}

type OrderRepository interface {
	Create(ctx context.Context, opts OrderCreateOpts) (*Order, error)
	FindAll(ctx context.Context, opts *OrderFindAllOpts) (*OrderList, error)
}
