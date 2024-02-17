package restaurantratingbiz

import (
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
	"errors"
)

type UserRatingRestaurantStore interface {
	Create(context context.Context, data *restaurantratingmodel.CreateRatingRestaurant) error
	Find(conntext context.Context, conditions map[string]interface{}) (*restaurantratingmodel.Rating, error)
}

type userRatingRestaurantBiz struct {
	store UserRatingRestaurantStore
}

func NewUserRatingRestaurantBiz(store UserRatingRestaurantStore) *userRatingRestaurantBiz {
	return &userRatingRestaurantBiz{store: store}
}

func (biz *userRatingRestaurantBiz) RatingRestaurant(context context.Context, data *restaurantratingmodel.CreateRatingRestaurant) error {
	_, err := biz.store.Find(context, map[string]interface{}{"point": data.Point})
	if err != nil {
		return errors.New("Bạn đã chấm điểm cho nhà hàng này rồi")
	}

	err = biz.store.Create(context, data)
	if err != nil {
		return err
	}
	return nil
}
