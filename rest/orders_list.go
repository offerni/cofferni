package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/offerni/cofferni"
)

func (srv *Server) OrdersList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	orders, err := srv.MenuService.OrderList(ctx)
	if err != nil {
		errorJSON, _ := json.Marshal(&ErrorResponse{
			Error: err.Error(),
		})

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJSON)
		return
	}

	ordersResponse := make([]OrderFetchResponse, len(orders.Orders))
	for i, order := range orders.Orders {
		ordersResponse[i] = OrderFetchResponse{
			CreatedAt:   order.CreatedAt,
			ID:          cofferni.OrderID(order.ID),
			ItemID:      order.ItemID,
			ModifiedAt:  order.ModifiedAt,
			Observation: order.Observation,
			Quantity:    order.Quantity,
		}
	}

	response := ListResponse[OrderFetchResponse]{
		Data: ordersResponse,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type OrderFetchResponse struct {
	CreatedAt   time.Time        `json:"created_at"`
	ID          cofferni.OrderID `json:"id"`
	ItemID      cofferni.ItemID  `json:"item_id"`
	ModifiedAt  time.Time        `json:"modified_at"`
	Observation *string          `json:"observation"`
	Quantity    uint             `json:"quantity"`
}
