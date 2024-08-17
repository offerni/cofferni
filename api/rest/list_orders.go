package rest

import (
	"encoding/json"
	"net/http"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/menu"
	"github.com/offerni/cofferni/utils"
)

func (srv *Server) ListOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	filterByFulfilled, err := utils.StringToBool(r.URL.Query().Get("fulfilled"))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	orders, err := srv.MenuService.OrderList(ctx, menu.OrderListOpts{
		FilterByFulfilled: filterByFulfilled,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ordersResponse := make([]OrderFetchResponse, len(orders.Orders))
	for i, order := range orders.Orders {
		ordersResponse[i] = OrderFetchResponse{
			CreatedAt:    utils.FormatTime(order.CreatedAt),
			CustomerName: order.CustomerName,
			Fulfilled:    order.Fulfilled,
			ID:           cofferni.OrderID(order.ID),
			ItemID:       order.ItemID,
			ItemName:     order.ItemName,
			ModifiedAt:   utils.FormatTime(order.ModifiedAt),
			Observation:  order.Observation,
			Quantity:     order.Quantity,
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
	CreatedAt    string           `json:"created_at"`
	CustomerName string           `json:"customer_name"`
	Fulfilled    bool             `json:"fulfilled"`
	ID           cofferni.OrderID `json:"id"`
	ItemID       cofferni.ItemID  `json:"item_id"`
	ItemName     string           `json:"item_name"`
	ModifiedAt   string           `json:"modified_at"`
	Observation  *string          `json:"observation"`
	Quantity     uint             `json:"quantity"`
}
