package menu

import (
	"context"

	"github.com/offerni/cofferni"
)

func (svc *Service) OrderUpdate(ctx context.Context, opts UpdateOrderOpts) (*UpdateOrderResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	order, err := svc.orderRepo.Update(ctx, cofferni.OrderUpdateOpts{
		Fulfilled:   opts.Fulfilled,
		ID:          opts.ID,
		Observation: opts.Observation,
		Quantity:    opts.Quantity,
	})

	if err != nil {
		return nil, err
	}

	item, err := svc.itemRepo.Find(ctx, order.ItemID)
	if err != nil {
		return nil, err
	}

	return &UpdateOrderResponse{
		&OrderFetchResponse{
			CreatedAt:    order.CreatedAt,
			CustomerName: order.CustomerName,
			Fulfilled:    order.Fulfilled,
			ID:           cofferni.OrderID(order.ID),
			ItemID:       order.ItemID,
			ItemName:     item.Name,
			ModifiedAt:   order.ModifiedAt,
			Observation:  order.Observation,
			Quantity:     order.Quantity,
		},
	}, nil
}

type UpdateOrderResponse struct {
	*OrderFetchResponse
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
