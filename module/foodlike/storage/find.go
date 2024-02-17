package foodlikestorage

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(context context.Context, conditions map[string]interface{}) (*foodlikemodel.Likefood, error) {
	db := s.db

	var like foodlikemodel.Likefood

	if err := db.Where(conditions).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

	}
	return &like, nil
}
