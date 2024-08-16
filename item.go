package cofferni

import (
	"context"
	"time"
)

type Item struct {
	ID         string
	Name       string
	Available  bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type ItemFindAllOpts struct {
}

type ItemList struct {
	Data []Item
	// pagintion later maybe
}

type ItemRepository interface {
	FindAll(ctx context.Context, opts ItemFindAllOpts) (*ItemList, error)
}
