package restaurantratingbiz

import (
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	"context"
	"errors"
)

type DeleteRatingRestaurantStore interface {
	Find(conntext context.Context, conditions map[string]interface{}) (*restaurantratingmodel.Rating, error)
	Delete(context context.Context, id int) error
}

type deleteRatingRestaurantBiz struct {
	store DeleteRatingRestaurantStore
}

func NewDeleteRatingRestaurantBiz(store DeleteRatingRestaurantStore) *deleteRatingRestaurantBiz {
	return &deleteRatingRestaurantBiz{store: store}
}

func (biz *deleteRatingRestaurantBiz) DeleteRatingRestaurant(context context.Context, id int) error {
	result, err := biz.store.Find(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if result.Comment == "" {
		return errors.New("comment deleted")
	}

	// if result != nil {
	// 	return errors.New("comment deleted")
	// }

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
