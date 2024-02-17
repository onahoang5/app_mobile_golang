package restaurantstorage

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

// tim db
func (s *sqlsStore) FindDataWithCondition(context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	var data restaurantmodel.Restaurant
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}
	return &data, nil
}

// func (s *sqlsStore) IsNameExistsInRestaurant(context context.Context, name string, restaurantId int) (bool, error) {
// 	var count int64
// 	if err := s.db.Model(&restaurantmodel.Restaurant{}).Where("name = ? AND restaurant_id = ?", name, restaurantId).Count(&count).Error; err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
