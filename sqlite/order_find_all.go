package sqlite

import (
	"context"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *orderRepo) FindAll(ctx context.Context) (*cofferni.OrderList, error) {
	result := []*models.Order{}

	err := repo.db.DB.WithContext(ctx).Find(&result).Error
	if err != nil {
		return nil, err
	}

	orders := make([]*cofferni.Order, len(result))
	for i, order := range result {
		orders[i] = &cofferni.Order{
			ID:           cofferni.OrderID(order.ID),
			ItemID:       cofferni.ItemID(order.ItemID),
			Observation:  order.Observation,
			Quantity:     order.Quantity,
			CreatedAt:    order.CreatedAt,
			ModifiedAt:   order.ModifiedAt,
			CustomerName: order.CustomerName,
			Fulfilled:    order.Fulfilled,
		}
	}

	return &cofferni.OrderList{Data: orders}, nil
}
