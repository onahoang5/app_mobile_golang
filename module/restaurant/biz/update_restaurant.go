package restaurantbiz

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(context context.Context, id int, data *restaurantmodel.RestaurantUpdate,
	) error
}

// interface khai báo nơi chúng ta dùng nó

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	// if err := data.Validate(); err != nil {
	// 	// return err
	// 	return common.ErrInvalidRequest(err)
	// }
	result, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		// return err
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	return nil
}
