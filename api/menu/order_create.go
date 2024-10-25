package menu

import (
	"context"

	"github.com/offerni/cofferni"
)

func (svc *Service) OrderCreate(ctx context.Context, opts CreateOrderOpts) (*CreateOrderResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	item, err := svc.itemRepo.Find(ctx, opts.ItemID)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, ErrItemNotFound
	}

	order, err := svc.orderRepo.Create(ctx, cofferni.OrderCreateOpts{
		CustomerName: opts.CustomerName,
		ItemID:       opts.ItemID,
		Observation:  opts.Observation,
		Quantity:     opts.Quantity,
	})

	if err != nil {
		return nil, err
	}

	return &CreateOrderResponse{
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

type CreateOrderOpts struct {
	CustomerName string
	ItemID       cofferni.ItemID
	Observation  *string
	Quantity     uint
}

type CreateOrderResponse struct {
	*OrderFetchResponse
}

func (opts CreateOrderOpts) Validate() error {
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
