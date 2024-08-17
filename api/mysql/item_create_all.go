package mysql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/mysql/models"
)

func (repo *itemRepo) CreateAll(ctx context.Context, opts cofferni.ItemCreateAllOpts) (*cofferni.ItemList, error) {
	items := make([]*models.Item, len(opts.Items))

	for i, item := range opts.Items {
		items[i] = &models.Item{
			Available:   true,
			CreatedAt:   time.Now(),
			Description: item.Description,
			ID:          uuid.New().String(),
			ModifiedAt:  time.Now(),
			Name:        item.Name,
		}
	}

	err := repo.db.DB.WithContext(ctx).Create(items).Error

	if err != nil {
		return nil, err
	}

	domainItems := make([]*cofferni.Item, len(items))
	for i, item := range items {
		domainItems[i] = &cofferni.Item{
			Available:  item.Available,
			CreatedAt:  item.CreatedAt,
			ID:         cofferni.ItemID(item.ID),
			ModifiedAt: item.ModifiedAt,
			Name:       item.Name,
		}
	}

	return &cofferni.ItemList{Data: domainItems}, nil
}
