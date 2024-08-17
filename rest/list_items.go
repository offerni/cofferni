package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/offerni/cofferni"
)

func (srv *Server) ListItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	items, err := srv.MenuService.ItemList(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	itemsResponse := make([]ItemFetchResponse, len(items.Items))
	for i, item := range items.Items {
		itemsResponse[i] = ItemFetchResponse{
			CreatedAt:  item.CreatedAt,
			ID:         cofferni.ItemID(item.ID),
			Available:  item.Available,
			ModifiedAt: item.ModifiedAt,
			Name:       item.Name,
		}
	}

	response := ListResponse[ItemFetchResponse]{
		Data: itemsResponse,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type ItemFetchResponse struct {
	Available  bool            `json:"available"`
	CreatedAt  time.Time       `json:"created_at"`
	ID         cofferni.ItemID `json:"id"`
	ModifiedAt time.Time       `json:"modified_at"`
	Name       string          `json:"name"`
}
