package foodlikestorage

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

func (s *sqlStore) GetFoodLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		FoodId     int `gorm:"column:restaurant_id"`
		LikedCount int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(foodlikemodel.Likefood{}.TableName()).Select("food_id, count(food_id) as count").
		Where("food_id in (?)", ids).
		Group("food_id").
		Find(&listLike).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.FoodId] = item.LikedCount
	}

	return result, nil
}

var timeLayout = "2006-01-02T15:04:05.99999"
var timeLayoutNoSecond = "2006-01-02T15:04:05"

func (s *sqlStore) GetUserLikeFood(context context.Context, conditions map[string]interface{}, filter *foodlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []foodlikemodel.Likefood
	db := s.db.Table(foodlikemodel.Likefood{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.FoodId > 0 {
			db = db.Where("food_id = ?", v.FoodId)
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
