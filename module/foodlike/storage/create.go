package foodlikestorage

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *foodlikemodel.Likefood) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
