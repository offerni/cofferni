package sqlite

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *orderRepo) Create(ctx context.Context, opts cofferni.OrderCreateOpts) (*cofferni.Order, error) {
	id := uuid.New().String()

	order := &models.Order{
		CreatedAt:   time.Now(),
		ID:          id,
		ItemID:      string(opts.ItemID),
		ModifiedAt:  time.Now(),
		Observation: opts.Observation,
		Quantity:    opts.Quantity,
		Fulfilled:   false,
	}

	err := repo.db.DB.WithContext(ctx).Create(order).Error

	if err != nil {
		return nil, err
	}

	return &cofferni.Order{
		CreatedAt:   order.CreatedAt,
		ID:          cofferni.OrderID(order.ID),
		ItemID:      cofferni.ItemID(order.ItemID),
		ModifiedAt:  order.ModifiedAt,
		Observation: order.Observation,
		Quantity:    order.Quantity,
		Fulfilled:   order.Fulfilled,
	}, nil
}
