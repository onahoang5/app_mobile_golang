package foodlikestorage

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, userId, foodId int) error {
	db := s.db

	if err := db.Table(foodlikemodel.Likefood{}.TableName()).Where("user_id = ? and food_id = ?", userId, foodId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
