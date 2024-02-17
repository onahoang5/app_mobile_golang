package restaurantratingrepo

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

type ListRatingRestaurantStore interface {
	ListDataRatingRestaurant(context context.Context, filter *restaurantratingmodel.Filter,
		paging *common.Paging, moreKeys ...string) ([]restaurantratingmodel.Rating, error)

	// GetUserRatingRestaurant(context context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter,
	// 	paging *common.Paging,
	// 	moreKeys ...string,
	// ) ([]common.SimpleUser, error)
}

type listRatingRestaurantRepo struct {
	store ListRatingRestaurantStore
}

func NewListRatingRestaurantRepo(store ListRatingRestaurantStore) *listRatingRestaurantRepo {
	return &listRatingRestaurantRepo{store: store}
}

func (biz *listRatingRestaurantRepo) ListRatingRestaurant(context context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging) ([]restaurantratingmodel.Rating, error) {
	result, err := biz.store.ListDataRatingRestaurant(context, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	return result, nil
}
