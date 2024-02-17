package restaurantstorage

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlsStore) UpdateRestaurant(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id= ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlsStore) IncreaseLikeCount(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlsStore) DecreaseLikeCount(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
