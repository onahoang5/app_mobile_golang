package restaurantlikestorage

import (
	"Food-delivery/common"
	restaurantlikemodel "Food-delivery/module/restaurantlike/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error) {

	db := s.db.Table(restaurantlikemodel.Like{}.TableName())

	var like restaurantlikemodel.Like

	if err := db.Where(conditions).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &like, nil
}
