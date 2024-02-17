package foodlikebiz

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"Food-delivery/pubsub"
	"context"
	"log"
)

type DisLikeFoodStore interface {
	Delete(context context.Context, userId, foodId int) error
	Find(context context.Context, conditions map[string]interface{}) (*foodlikemodel.Likefood, error)
}

// type DecreaseLikeCountFoodStore interface {
// 	DecreaseLikeCountFood(context context.Context, id int) error
// }

type userDisLikeFoodBiz struct {
	store DisLikeFoodStore
	// Decrstore DecreaseLikeCountFoodStore
	ps pubsub.Pubsub
}

func NewDisLikeFoodBiz(store DisLikeFoodStore, ps pubsub.Pubsub) *userDisLikeFoodBiz {
	return &userDisLikeFoodBiz{store: store, ps: ps}
}

func (biz *userDisLikeFoodBiz) DisLikeFood(context context.Context, userId, foodId int) error {
	_, err := biz.store.Find(context, map[string]interface{}{"user_id": userId, "food_id": foodId})

	if err != nil {
		return foodlikemodel.ErrAlreadyLikedFood()
	}

	err = biz.store.Delete(context, userId, foodId)
	if err != nil {
		return foodlikemodel.ErrCannotLikeFood(err)
	}

	// err = biz.Decrstore.DecreaseLikeCountFood(context, foodId)
	// if err != nil {
	// 	log.Println(err)
	// }

	if err = biz.ps.Publish(context, common.TopicUserDisLikeFood,
		pubsub.NewMessage(&foodlikemodel.Likefood{
			FoodId: foodId,
			UserId: userId,
		})); err != nil {
		log.Println(err)
	}

	return nil
}
