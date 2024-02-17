package foodbiz

import (
	foodmodel "Food-delivery/module/food/model"
	"context"
	"errors"
)

type UpdateFoodStore interface {
	Find(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*foodmodel.Food, error)
	UpdateData(
		context context.Context,
		id int,
		data *foodmodel.FoodUpdate,
	) error
}

type updateFoodBiz struct {
	store UpdateFoodStore
	// requester common.Requester
}

func NewUpdateFoodBiz(store UpdateFoodStore) *updateFoodBiz {
	return &updateFoodBiz{store: store}
}

func (biz *updateFoodBiz) UpdateFood(context context.Context, id int, data *foodmodel.FoodUpdate) error {
	result, err := biz.store.Find(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(context, id, data); err != nil {
		return err
	}
	return nil
}
