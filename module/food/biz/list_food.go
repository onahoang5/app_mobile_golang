package foodbiz

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

type ListFoodRepo interface {
	ListFood(
		context context.Context,
		filter *foodmodel.Filter,
		paging *common.Paging,
	) ([]foodmodel.Food, error)
}

type listFoodBiz struct {
	repo ListFoodRepo
}

func NewListFoodtBiz(repo ListFoodRepo) *listFoodBiz {
	return &listFoodBiz{repo: repo}
}

func (biz *listFoodBiz) ListFood(
	context context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
) ([]foodmodel.Food, error) {

	result, err := biz.repo.ListFood(context, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
