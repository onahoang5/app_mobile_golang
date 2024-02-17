package foodstorage

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

func (s *sqlStore) DeleteFoodSoft(context context.Context, id int) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
