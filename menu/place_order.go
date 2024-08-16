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
		ItemID:      opts.ItemID,
		Observation: opts.Observation,
		Quantity:    opts.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return &PlaceOrderResponse{
		CreatedAt:   order.CreatedAt,
		ID:          cofferni.OrderID(order.ID),
		ItemID:      string(order.ItemID),
		ModifiedAt:  order.ModifiedAt,
		Observation: order.Observation,
		Quantity:    order.Quantity,
	}, nil
}

type PlaceOrderOpts struct {
	ItemID      cofferni.ItemID
	Observation *string
	Quantity    uint
}

type PlaceOrderResponse struct {
	CreatedAt   time.Time
	ID          cofferni.OrderID
	ItemID      string
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

func (opts PlaceOrderOpts) Validate() error {
	if opts.ItemID == "" {
		return ErrItemIDIsRequired
	}

	if opts.Quantity == 0 {
		return ErrQuantityIsRequired
	}

	return nil
}
