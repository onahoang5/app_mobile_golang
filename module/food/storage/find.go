package foodstorage

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error) {

	db := s.db.Table(foodmodel.Food{}.TableName())

	var data foodmodel.Food

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}

func (s *sqlStore) IsNameExistsInRestaurant(context context.Context, name string, restaurantId int) (bool, error) {
	var count int64
	if err := s.db.Model(&foodmodel.Food{}).Where("name = ? AND restaurant_id = ?", name, restaurantId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
