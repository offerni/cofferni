package sqlite

import (
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/connection"
)

type orderRepo struct {
	db *connection.DB
}

func NewOrderRepository(db *connection.DB) cofferni.OrderRepository {
	return &orderRepo{db: db}
}
