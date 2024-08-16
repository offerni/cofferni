package sqlite

import (
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/connection"
)

type itemRepo struct {
	db *connection.DB
}

func NewItemRepository(db *connection.DB) cofferni.ItemRepository {
	return &itemRepo{db: db}
}
