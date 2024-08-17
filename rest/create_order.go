package rest

import (
	"encoding/json"
	"net/http"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/menu"
)

func (srv *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var request CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorJSON, _ := json.Marshal(&ErrorResponse{
			Error: "Invalid Body",
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJSON)
		return
	}

	order, err := srv.MenuService.OrderCreate(ctx, menu.CreateOrderOpts{
		CustomerName: request.CustomerName,
		ItemID:       request.ItemID,
		Observation:  request.Observation,
		Quantity:     request.Quantity,
	})
	if err != nil {
		errorJSON, _ := json.Marshal(&ErrorResponse{
			Error: err.Error(),
		})

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorJSON)
		return
	}

	response := CreateOrderResponse{
		&OrderFetchResponse{
			CreatedAt:   order.CreatedAt,
			ID:          order.ID,
			ItemID:      order.ItemID,
			ItemName:    order.ItemName,
			ModifiedAt:  order.ModifiedAt,
			Observation: order.Observation,
			Quantity:    order.Quantity,
		},
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type CreateOrderRequest struct {
	CustomerName string          `json:"customer_name"`
	ItemID       cofferni.ItemID `json:"item_id"`
	Observation  *string         `json:"observation"`
	Quantity     uint            `json:"quantity"`
}

type CreateOrderResponse struct {
	*OrderFetchResponse
}
