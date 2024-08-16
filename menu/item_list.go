package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) ItemList(ctx context.Context) (*ItemListResponse, error) {
	items, err := svc.ItemRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	itemsResponse := make([]*ItemFetchResponse, len(items.Data))

	for i, item := range items.Data {
		itemsResponse[i] = &ItemFetchResponse{
			Available:  item.Available,
			CreatedAt:  item.CreatedAt,
			ID:         cofferni.ItemID(item.ID),
			ModifiedAt: item.ModifiedAt,
			Name:       item.Name,
		}
	}

	return &ItemListResponse{
		Items: itemsResponse,
	}, nil
}

type ItemFetchResponse struct {
	Available  bool
	CreatedAt  time.Time
	ID         cofferni.ItemID
	ModifiedAt time.Time
	Name       string
}

type ItemListResponse struct {
	Items []*ItemFetchResponse
}
