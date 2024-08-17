package cofferni

import (
	"context"
	"time"
)

type OrderID string

type Order struct {
	CreatedAt    time.Time
	CustomerName string
	Fulfilled    bool
	ID           OrderID
	ItemID       ItemID
	ModifiedAt   time.Time
	Observation  *string
	Quantity     uint
}

type OrderFindAllOpts struct {
	Fulfilled *bool
}

type OrderList struct {
	Data []*Order
	// pagintion later maybe
}

type OrderCreateOpts struct {
	CustomerName string
	ItemID       ItemID
	Observation  *string
	Quantity     uint
}

func (opts OrderCreateOpts) Validate() error {
	if opts.CustomerName == "" {
		return ErrCustomerNameIsRequired
	}

	if opts.ItemID == "" {
		return ErrItemIDIsRequired
	}

	if opts.Quantity == 0 {
		return ErrQuantityIsRequired
	}

	return nil
}

type OrderUpdateOpts struct {
	Fulfilled   *bool
	ID          OrderID
	Observation *string
	Quantity    *uint
}

func (opts OrderUpdateOpts) Validate() error {
	if opts.ID == "" {
		return ErrIDIsRequired
	}

	if opts.Quantity != nil && *opts.Quantity == 0 {
		return ErrQuantityIsRequired
	}

	if opts.Observation != nil && *opts.Observation == "" {
		return ErrObservationIsRequired
	}

	if opts.Fulfilled == nil {
		return ErrFulfilledIsRequired
	}

	return nil
}

type OrderRepository interface {
	Create(ctx context.Context, opts OrderCreateOpts) (*Order, error)
	FindAll(ctx context.Context, opts OrderFindAllOpts) (*OrderList, error)
	Update(ctx context.Context, opts OrderUpdateOpts) (*Order, error)
}
