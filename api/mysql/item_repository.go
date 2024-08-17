package mysql

import (
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/mysql/connection"
)

type itemRepo struct {
	db *connection.DB
}

func NewItemRepository(db *connection.DB) cofferni.ItemRepository {
	return &itemRepo{db: db}
}
