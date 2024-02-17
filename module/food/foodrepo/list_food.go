package foodrepo

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

type ListFoodStore interface {
	ListDataFoodByconditons(context context.Context, filter *foodmodel.Filter, paging *common.Paging,
		moreKeys ...string) ([]foodmodel.Food, error)
}

type listFoodRepo struct {
	store ListFoodStore
}

func NewListFoodRepo(store ListFoodStore) *listFoodRepo {
	return &listFoodRepo{store: store}
}

func (biz *listFoodRepo) ListFood(context context.Context,
	filter *foodmodel.Filter, paging *common.Paging) ([]foodmodel.Food, error) {
	result, err := biz.store.ListDataFoodByconditons(context, filter, paging, "Restaurant")
	if err != nil {
		return nil, err
	}

	return result, nil
}
