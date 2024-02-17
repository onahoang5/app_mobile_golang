package restaurantratingrepo

import (
	"Food-delivery/common"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
)

type GetRatingRestaurantStore interface {
	GetUserRatingRestaurant(context context.Context, conditions map[string]interface{}, filter *restaurantratingmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type getRatingRestaurantRepo struct {
	store GetRatingRestaurantStore
}

func NewGetRatingRestaurantRepo(store GetRatingRestaurantStore) *getRatingRestaurantRepo {
	return &getRatingRestaurantRepo{store: store}
}

func (biz *getRatingRestaurantRepo) GetRatingRestaurant(
	context context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUserRatingRestaurant(context, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}
	return users, nil
}
