package restaurantrepository

import (
	"Food-delivery/common"
	restaurantmodel "Food-delivery/module/restaurant/model"
	"context"
)

type GetRestaurantStore interface {
	FindDataWithCondition(context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantRepo struct {
	store     GetRestaurantStore
	requester common.Requester
}

func NewGetRestaurantRepo(store GetRestaurantStore, requester common.Requester) *getRestaurantRepo {
	return &getRestaurantRepo{store: store, requester: requester}
}

func (biz *getRestaurantRepo) GetRestaurant(
	context context.Context,
	id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id}, "User")
	if err != nil {
		return nil, err
	}

	return result, nil
}
