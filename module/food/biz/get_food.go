package foodbiz

import (
	"Food-delivery/common"
	foodmodel "Food-delivery/module/food/model"
	"context"
)

type FindFoodRepo interface {
	GetFood(context context.Context, id int) (*foodmodel.Food, error)
}

type getFoodBiz struct {
	repo FindFoodRepo
}

func NewGetFoodBiz(repo FindFoodRepo) *getFoodBiz {
	return &getFoodBiz{repo: repo}
}

func (biz *getFoodBiz) GetFood(context context.Context, id int) (*foodmodel.Food, error) {

	result, err := biz.repo.GetFood(context, id)

	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return result, nil
}
