package foodlikebiz

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"Food-delivery/pubsub"
	"context"
	"log"
)

type UserLikeFoodStore interface {
	Find(context context.Context, conditions map[string]interface{}) (*foodlikemodel.Likefood, error)
	Create(context context.Context, data *foodlikemodel.Likefood) error
}

// type IncreaseLikeCountFoodStore interface {
// 	IncreaseLikeCountFood(context context.Context, id int) error
// }

type userLikeFoodBiz struct {
	store UserLikeFoodStore
	// increStore IncreaseLikeCountFoodStore
	ps pubsub.Pubsub
}

func NewUserLikeFoodBiz(store UserLikeFoodStore, ps pubsub.Pubsub) *userLikeFoodBiz {
	return &userLikeFoodBiz{store: store, ps: ps}
}

func (biz *userLikeFoodBiz) LikeFood(context context.Context, data *foodlikemodel.Likefood) error {
	_, err := biz.store.Find(context, map[string]interface{}{"user_id": data.UserId, "food_id": data.FoodId})

	if err == nil {
		return foodlikemodel.ErrAlreadyLikedFood()
	}

	err = biz.store.Create(context, data)

	if err != nil {
		return foodlikemodel.ErrCannotLikeFood(err)
	}
	// if err = biz.increStore.IncreaseLikeCountFood(context, data.FoodId); err != nil {
	// 	log.Println(err)
	// }

	if err = biz.ps.Publish(context, common.TopicUserLikeFood, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
