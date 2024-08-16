package cofferni

import (
	"context"
	"time"
)

type OrderID string

type Order struct {
	CreatedAt   time.Time
	Fulfilled   bool
	ID          OrderID
	ItemID      ItemID
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

type OrderCreateOpts struct {
	ItemID      ItemID
	Observation *string
	Quantity    uint
}

type OrderFindAllOpts struct {
}

type OrderList struct {
	Data []*Order
	// pagintion later maybe
}

type OrderRepository interface {
	Create(ctx context.Context, opts OrderCreateOpts) (*Order, error)
	FindAll(ctx context.Context) (*OrderList, error)
}

func (opts OrderCreateOpts) Validate() error {
	if opts.ItemID == "" {
		return ErrItemIDIsRequired
	}

	if opts.Quantity == 0 {
		return ErrQuantityIsRequired
	}

	return nil
}
