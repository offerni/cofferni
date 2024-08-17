package cofferni

import (
	"context"
	"time"
)

type ItemID string

type Item struct {
	Available   bool
	CreatedAt   time.Time
	Description *string
	ID          ItemID
	ModifiedAt  time.Time
	Name        string
}

type ItemFindAllOpts struct {
	Available *bool
}

type ItemList struct {
	Data []*Item
	// pagintion later maybe
}

type ItemCreateOpts struct {
	Available   bool
	Description *string
	Name        string
}

type ItemCreateAllOpts struct {
	Items []*ItemCreateOpts
}

type ItemRepository interface {
	FindAll(ctx context.Context, opts ItemFindAllOpts) (*ItemList, error)
	Find(ctx context.Context, ID ItemID) (*Item, error)
	CreateAll(ctx context.Context, opts ItemCreateAllOpts) (*ItemList, error)
}

// TODO: Add Validation
