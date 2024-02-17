package restaurantratingstorage

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

func (s *sqlStore) ListDataRatingRestaurant(context context.Context, filter *restaurantratingmodel.Filter,
	paging *common.Paging, moreKeys ...string) ([]restaurantratingmodel.Rating, error) {
	var result []restaurantratingmodel.Rating
	db := s.db.Table(restaurantratingmodel.Rating{}.TableName()).Where("status in(1)")

	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
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

	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func (s *sqlStore) GetRestaurantRating(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int     `gorm:"column:restaurant_id"`
		Comment      int     `gorm:"column:count"`
		Point        float64 `gorm:"column:point"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantratingmodel.Rating{}.TableName()).Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.Comment
	}

	return result, nil
}

var timeLayout = "2006-01-02T15:04:05.99999"
var timeLayoutNoSecond = "2006-01-02T15:04:05"

func (s *sqlStore) GetUserRatingRestaurant(context context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []restaurantratingmodel.Rating
	db := s.db.Table(restaurantratingmodel.Rating{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format(timeLayoutNoSecond))
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))
	}

	if err := db.Limit(int(paging.Limit)).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User // when len of two arr not equal (mean: user may is null) => system crashing
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
