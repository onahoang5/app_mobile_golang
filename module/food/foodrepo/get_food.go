package foodrepo

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

type GetFoodStore interface {
	Find(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error)
}

type getFoodRepo struct {
	store GetFoodStore
}

func NewGetFoodRepo(store GetFoodStore) *getFoodRepo {
	return &getFoodRepo{store: store}
}

func (biz *getFoodRepo) GetFood(context context.Context, id int) (*foodmodel.Food, error) {

	result, err := biz.store.Find(context, map[string]interface{}{"id": id}, "Restaurant")
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
