package restaurantbiz

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"
)

type GetRestaurantRepo interface {
	GetRestaurant(
		context context.Context,
		id int) (*restaurantmodel.Restaurant, error)
}

// interface khai báo nơi chúng ta dùng nó

type getRestaurantBiz struct {
	repo GetRestaurantRepo
}

func NewGetRestaurantBiz(repo GetRestaurantRepo) *getRestaurantBiz {
	return &getRestaurantBiz{repo: repo}
}

func (biz *getRestaurantBiz) GetRestaurant(
	context context.Context,
	id int) (*restaurantmodel.Restaurant, error) {

	result, err := biz.repo.GetRestaurant(context, id)

	if err != nil {

		if err != nil {
			return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
		}

	}

	return result, err
}
