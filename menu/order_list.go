package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) OrderList(ctx context.Context) (*OrderListResponse, error) {
	orders, err := svc.orderRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	items, err := svc.itemRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	itemMap := make(map[cofferni.ItemID]*cofferni.Item, len(items.Data))
	for _, item := range items.Data {
		itemMap[item.ID] = item
	}

	ordersResponse := make([]*OrderFetchResponse, len(orders.Data))
	for i, order := range orders.Data {
		ordersResponse[i] = &OrderFetchResponse{
			CreatedAt:   order.CreatedAt,
			ID:          cofferni.OrderID(order.ID),
			ItemID:      cofferni.ItemID(order.ItemID),
			ItemName:    itemMap[cofferni.ItemID(order.ItemID)].Name,
			ModifiedAt:  order.ModifiedAt,
			Observation: order.Observation,
			Quantity:    order.Quantity,
		}
	}

	return &OrderListResponse{
		Orders: ordersResponse,
	}, nil
}

type OrderFetchResponse struct {
	CreatedAt   time.Time
	ID          cofferni.OrderID
	ItemID      cofferni.ItemID
	ItemName    string
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

type OrderListResponse struct {
	Orders []*OrderFetchResponse
}
