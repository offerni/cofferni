package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) UpdateOrder(ctx context.Context, opts UpdateOrderOpts) (*UpdateOrderResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	order, err := svc.OrderRepo.Update(ctx, cofferni.OrderUpdateOpts{
		Fulfilled:   opts.Fulfilled,
		ID:          opts.ID,
		Observation: opts.Observation,
		Quantity:    opts.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return &UpdateOrderResponse{
		CreatedAt:   order.CreatedAt,
		ID:          cofferni.OrderID(order.ID),
		ItemID:      order.ItemID,
		ModifiedAt:  order.ModifiedAt,
		Observation: order.Observation,
		Quantity:    order.Quantity,
	}, nil
}

type UpdateOrderResponse struct {
	CreatedAt   time.Time
	ID          cofferni.OrderID
	ItemID      cofferni.ItemID
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

type UpdateOrderOpts struct {
	Fulfilled   *bool
	ID          cofferni.OrderID
	Observation *string
	Quantity    *uint
}

func (opts UpdateOrderOpts) Validate() error {
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
