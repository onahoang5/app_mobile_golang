package foodbiz

import (
	foodmodel "Food-delivery/module/food/model"
	"context"
	"errors"
	"log"
)

type CreateFoodStore interface {
	CreateFood(context context.Context, data *foodmodel.FoodCreate) error
	IsNameExistsInRestaurant(context context.Context, name string, restaurantId int) (bool, error)
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*foodmodel.Food, error)
}

// interface khai báo nơi chúng ta dùng nó

type createFoodBiz struct {
	store CreateFoodStore
}

func NewCreateFoodtBiz(store CreateFoodStore) *createFoodBiz {
	return &createFoodBiz{store: store}
}

func (biz *createFoodBiz) CreateFood(context context.Context, data *foodmodel.FoodCreate) error {

	result, err := biz.store.IsNameExistsInRestaurant(context, data.Name, data.RestaurantId)

	if err != nil {
		return err
	}

	if result {
		return errors.New("Name already exists in the restaurant")
	}

	_, err = biz.store.Find(context, map[string]interface{}{"id": data.Id})

	if err == nil {
		return foodmodel.ErrAlreadyFoodRestaurant()
	}

	err = biz.store.CreateFood(context, data)

	if err != nil {
		log.Println(err)
	}

	return nil
}
