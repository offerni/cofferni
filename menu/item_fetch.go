package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) ItemFetch(ctx context.Context, ID cofferni.ItemID) (*ItemFetchResponse, error) {
	item, err := svc.itemRepo.Find(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &ItemFetchResponse{
		Available:   item.Available,
		CreatedAt:   item.CreatedAt,
		Description: item.Description,
		ID:          cofferni.ItemID(item.ID),
		ModifiedAt:  item.ModifiedAt,
		Name:        item.Name,
	}, nil
}

type ItemFetchResponse struct {
	Available   bool
	CreatedAt   time.Time
	Description string
	ID          cofferni.ItemID
	ModifiedAt  time.Time
	Name        string
}
