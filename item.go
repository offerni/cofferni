package cofferni

import (
	"context"
	"time"
)

type ItemID string

type Item struct {
	Available  bool
	CreatedAt  time.Time
	ID         ItemID
	ModifiedAt time.Time
	Name       string
}

type ItemFindAllOpts struct {
}

type ItemList struct {
	Data []*Item
	// pagintion later maybe
}

type ItemRepository interface {
	FindAll(ctx context.Context) (*ItemList, error)
}
