package sqlite

import (
	"context"

	"github.com/offerni/cofferni"
	"github.com/offerni/cofferni/sqlite/models"
)

func (repo *itemRepo) Find(ctx context.Context, ID cofferni.ItemID) (*cofferni.Item, error) {
	result := models.Item{}

	query := repo.db.DB.WithContext(ctx)

	err := query.Where("id = ?", ID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &cofferni.Item{
		Available:   result.Available,
		CreatedAt:   result.CreatedAt,
		Description: result.Description,
		ID:          cofferni.ItemID(result.ID),
		ModifiedAt:  result.ModifiedAt,
		Name:        result.Name,
	}, nil
}
