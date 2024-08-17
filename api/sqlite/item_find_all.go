package sqlite

import (
	"context"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *itemRepo) FindAll(ctx context.Context, opts cofferni.ItemFindAllOpts) (*cofferni.ItemList, error) {
	result := []*models.Item{}

	query := repo.db.DB.WithContext(ctx)

	if opts.Available != nil {
		query = query.Where("available = ?", *opts.Available)
	}

	err := query.Order("name").Find(&result).Error
	if err != nil {
		return nil, err
	}

	items := make([]*cofferni.Item, len(result))
	for i, item := range result {
		items[i] = &cofferni.Item{
			ID:          cofferni.ItemID(item.ID),
			Name:        item.Name,
			Available:   item.Available,
			CreatedAt:   item.CreatedAt,
			ModifiedAt:  item.ModifiedAt,
			Description: item.Description,
		}
	}

	return &cofferni.ItemList{Data: items}, nil
}
