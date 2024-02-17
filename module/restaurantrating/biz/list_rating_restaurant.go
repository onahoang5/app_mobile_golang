package restaurantratingbiz

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

type ListRatingRestaurantRepo interface {
	ListRatingRestaurant(
		context context.Context,
		filter *restaurantratingmodel.Filter,
		paging *common.Paging,
	) ([]restaurantratingmodel.Rating, error)
}

type listRatingRestaurantBiz struct {
	repo ListRatingRestaurantRepo
}

func NewListRatingRestaurantBiz(repo ListRatingRestaurantRepo) *listRatingRestaurantBiz {
	return &listRatingRestaurantBiz{repo: repo}
}

func (biz *listRatingRestaurantBiz) ListRatingRestaurant(
	context context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
) ([]restaurantratingmodel.Rating, error) {

	result, err := biz.repo.ListRatingRestaurant(context, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantratingmodel.EntityName, err)
	}

	return result, nil
}
