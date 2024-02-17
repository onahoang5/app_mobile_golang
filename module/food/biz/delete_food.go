package foodbiz

import (
	foodmodel "Food-delivery/module/food/model"
	"context"
	"errors"
)

type DeleteFoodStore interface {
	Find(context context.Context, conditions map[string]interface{}, moreKey ...string) (*foodmodel.Food, error)
	DeleteFoodSoft(context context.Context, id int) error
}

type deleteFoodBiz struct {
	store DeleteFoodStore
}

func NewDeleteFoodBiz(store DeleteFoodStore) *deleteFoodBiz {
	return &deleteFoodBiz{store: store}
}

func (biz *deleteFoodBiz) DeleteFood(context context.Context, id int) error {
	result, err := biz.store.Find(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.DeleteFoodSoft(context, id); err != nil {
		return err
	}

	return nil
}
