package sqlite

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *itemRepo) CreateAll(ctx context.Context, opts cofferni.ItemCreateAllOpts) (*cofferni.ItemList, error) {

	items := make([]*models.Item, len(opts.Items))
	for i, item := range opts.Items {
		items[i] = &models.Item{
			CreatedAt:  time.Now(),
			ID:         uuid.New().String(),
			ModifiedAt: time.Now(),
			Name:       item.Name,
			Available:  true,
		}
	}

	err := repo.db.DB.WithContext(ctx).Create(items).Error

	if err != nil {
		return nil, err
	}

	domainItems := make([]*cofferni.Item, len(items))
	for i, item := range items {
		domainItems[i] = &cofferni.Item{
			CreatedAt:  item.CreatedAt,
			ID:         cofferni.ItemID(item.ID),
			ModifiedAt: item.ModifiedAt,
			Name:       item.Name,
			Available:  item.Available,
		}
	}

	return &cofferni.ItemList{Data: domainItems}, nil
}
