package mysql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/mysql/models"
)

func (repo *orderRepo) Create(ctx context.Context, opts cofferni.OrderCreateOpts) (*cofferni.Order, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	id := uuid.New().String()

	order := &models.Order{
		CreatedAt:    time.Now(),
		CustomerName: opts.CustomerName,
		ID:           id,
		ItemID:       string(opts.ItemID),
		ModifiedAt:   time.Now(),
		Observation:  opts.Observation,
		Quantity:     opts.Quantity,
	}

	err := repo.db.DB.WithContext(ctx).FirstOrCreate(order).Error
	if err != nil {
		return nil, err
	}

	return &cofferni.Order{
		CreatedAt:    order.CreatedAt,
		CustomerName: order.CustomerName,
		Fulfilled:    order.Fulfilled,
		ID:           cofferni.OrderID(order.ID),
		ItemID:       cofferni.ItemID(order.ItemID),
		ModifiedAt:   order.ModifiedAt,
		Observation:  order.Observation,
		Quantity:     order.Quantity,
	}, nil
}
