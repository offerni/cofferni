package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) PlaceOrder(ctx context.Context, opts PlaceOrderOpts) (*PlaceOrderResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	order, err := svc.OrderRepo.Create(ctx, cofferni.OrderCreateOpts{
		CustomerName: opts.CustomerName,
		ItemID:       opts.ItemID,
		Observation:  opts.Observation,
		Quantity:     opts.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return &PlaceOrderResponse{
		CreatedAt:    order.CreatedAt,
		CustomerName: order.CustomerName,
		ID:           cofferni.OrderID(order.ID),
		ItemID:       order.ItemID,
		ModifiedAt:   order.ModifiedAt,
		Observation:  order.Observation,
		Quantity:     order.Quantity,
	}, nil
}

type PlaceOrderOpts struct {
	CustomerName string
	ItemID       cofferni.ItemID
	Observation  *string
	Quantity     uint
}

type PlaceOrderResponse struct {
	CreatedAt    time.Time
	CustomerName string
	ID           cofferni.OrderID
	ItemID       cofferni.ItemID
	ModifiedAt   time.Time
	Observation  *string
	Quantity     uint
}

func (opts PlaceOrderOpts) Validate() error {
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
