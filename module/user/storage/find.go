package userstorage

import (
	"Food-delivery/common"
	usermodel "Food-delivery/module/user/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(context context.Context,
	condition map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
