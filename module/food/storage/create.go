package foodstorage

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

func (s *sqlStore) CreateFood(context context.Context, data *foodmodel.FoodCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		// return err
		return common.ErrDB(err)
	}
	return nil
}
