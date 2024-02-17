package restaurantratingstorage

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(conntext context.Context, conditions map[string]interface{}) (*restaurantratingmodel.Rating, error) {
	db := s.db.Table(restaurantratingmodel.Rating{}.TableName())
	var rating restaurantratingmodel.Rating
	if err := db.Where(conditions).First(&rating).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &rating, nil
}
