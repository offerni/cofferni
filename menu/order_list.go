package menu

import (
	"context"
	"time"

	"github.com/offerni/cofferni"
)

func (svc *Service) OrderList(ctx context.Context) (*OrderListResponse, error) {
	orders, err := svc.OrderRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	ordersResponse := make([]*OrderFetchResponse, len(orders.Data))

	for i, order := range orders.Data {
		ordersResponse[i] = &OrderFetchResponse{
			CreatedAt:   order.CreatedAt,
			ID:          cofferni.OrderID(order.ID),
			ItemID:      cofferni.ItemID(order.ItemID),
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
	ModifiedAt  time.Time
	Observation *string
	Quantity    uint
}

type OrderListResponse struct {
	Orders []*OrderFetchResponse
}
