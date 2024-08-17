package rest

import (
	"encoding/json"
	"net/http"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/menu"
	"github.com/offerni/cofferni/utils"
)

func (srv *Server) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var request UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorJSON, _ := json.Marshal(&ErrorResponse{
			Error: "Invalid Body",
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJSON)
		return
	}

	order, err := srv.MenuService.OrderUpdate(ctx, menu.UpdateOrderOpts{
		Fulfilled:   request.Fulfilled,
		ID:          request.ID,
		Observation: request.Observation,
		Quantity:    request.Quantity,
	})
	if err != nil {
		errorJSON, _ := json.Marshal(&ErrorResponse{
			Error: err.Error(),
		})

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorJSON)
		return
	}

	response := UpdateOrderResponse{
		&OrderFetchResponse{
			CreatedAt:    utils.FormatTime(order.CreatedAt),
			CustomerName: order.CustomerName,
			Fulfilled:    order.Fulfilled,
			ID:           order.ID,
			ItemID:       order.ItemID,
			ItemName:     order.ItemName,
			ModifiedAt:   utils.FormatTime(order.CreatedAt),
			Observation:  order.Observation,
			Quantity:     order.Quantity,
		},
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type UpdateOrderRequest struct {
	Fulfilled   *bool            `json:"fulfilled"`
	ID          cofferni.OrderID `json:"id"`
	Observation *string          `json:"observation"`
	Quantity    *uint            `json:"quantity"`
}

type UpdateOrderResponse struct {
	*OrderFetchResponse
}
