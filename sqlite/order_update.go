package sqlite

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *orderRepo) Update(ctx context.Context, opts cofferni.OrderUpdateOpts) (*cofferni.Order, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	order := &models.Order{}

	err := repo.db.DB.WithContext(ctx).First(order, opts.ID).Error
	if err != nil {
		return nil, err
	}

	if opts.Fulfilled != nil {
		order.Fulfilled = *opts.Fulfilled
	}
	if opts.Observation != nil && *opts.Observation != "" {
		order.Observation = opts.Observation
	}
	if opts.Quantity != nil {
		order.Quantity = *opts.Quantity
	}

	order.ModifiedAt = time.Now()

	err = repo.db.DB.WithContext(ctx).Save(order).Error

	if err != nil {
		return nil, err
	}

	var orderResponse = models.Order{}

	repo.db.DB.Find(&orderResponse)

	return &cofferni.Order{
		CreatedAt:   orderResponse.CreatedAt,
		ID:          cofferni.OrderID(orderResponse.ID),
		ItemID:      cofferni.ItemID(orderResponse.ItemID),
		ModifiedAt:  orderResponse.ModifiedAt,
		Observation: orderResponse.Observation,
		Quantity:    orderResponse.Quantity,
		Fulfilled:   orderResponse.Fulfilled,
	}, nil
}
