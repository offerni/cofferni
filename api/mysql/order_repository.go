package mysql

import (
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/mysql/connection"
)

type orderRepo struct {
	db *connection.DB
}

func NewOrderRepository(db *connection.DB) cofferni.OrderRepository {
	return &orderRepo{db: db}
}
