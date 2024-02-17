package restaurantratingbiz

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

type GetRatingRestaurantRepo interface {
	GetRatingRestaurant(
		context context.Context,
		filter *restaurantratingmodel.Filter,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type getRatingRestaurantBiz struct {
	repo GetRatingRestaurantRepo
}

func NewGetRatingRestaurantBiz(repo GetRatingRestaurantRepo) *getRatingRestaurantBiz {
	return &getRatingRestaurantBiz{repo: repo}
}

func (biz *getRatingRestaurantBiz) GetRatingRestaurant(
	context context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {

	result, err := biz.repo.GetRatingRestaurant(context, filter, paging)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}

	return result, nil
}
