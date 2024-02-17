package restaurantratingstorage

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantratingmodel.CreateRatingRestaurant) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
