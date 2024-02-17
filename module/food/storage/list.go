package foodstorage

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

func (s *sqlStore) ListDataFoodByconditons(context context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]foodmodel.Food, error) {
	var result []foodmodel.Food

	db := s.db.Table(foodmodel.Food{}.TableName()).Where("status in (1)")

	if v := filter; v != nil {
		if v.CategoryId > 0 {
			db = db.Where("category_id = ?", v.CategoryId)
		}

		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}

		if len(v.Status) > 0 {
			db = db.Where("status in (?)", v.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))

	}

	//log.Fatalln("paigngggggg", paging)
	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
