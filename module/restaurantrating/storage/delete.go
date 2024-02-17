package restaurantratingstorage

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	db := s.db
	if err := db.Table(restaurantratingmodel.Rating{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"comment": "",
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
