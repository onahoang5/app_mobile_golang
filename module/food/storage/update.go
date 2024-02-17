package foodstorage

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(
	context context.Context,
	id int,
	data *foodmodel.FoodUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCountFood(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCountFood(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(foodmodel.Food{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
