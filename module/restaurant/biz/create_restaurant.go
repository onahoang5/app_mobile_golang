package restaurantbiz

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"
	"errors"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
	FindDataWithCondition(context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

// interface khai báo nơi chúng ta dùng nó

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {

	result, _ := biz.store.FindDataWithCondition(context, map[string]interface{}{"name": data.Name})

	if result != nil {
		return errors.New("Name restaurant already exists  ")
	}
	if err := data.Validate(); err != nil {
		// return err
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.CreateRestaurant(context, data); err != nil {
		// return err
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
