package menu

import (
	"context"

	"github.com/offerni/cofferni"
)

func (svc *Service) ItemList(ctx context.Context) (*ItemListResponse, error) {
	items, err := svc.itemRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	itemsResponse := make([]*ItemFetchResponse, len(items.Data))

	for i, item := range items.Data {
		itemsResponse[i] = &ItemFetchResponse{
			Available:   item.Available,
			CreatedAt:   item.CreatedAt,
			Description: item.Description,
			ID:          cofferni.ItemID(item.ID),
			ModifiedAt:  item.ModifiedAt,
			Name:        item.Name,
		}
	}

	return &ItemListResponse{
		Items: itemsResponse,
	}, nil
}

type ItemListResponse struct {
	Items []*ItemFetchResponse
}