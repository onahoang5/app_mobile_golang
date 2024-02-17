package restautantlikebiz

import (
	"Food-delivery/common"
	restaurantlikemodel "Food-delivery/module/restaurantlike/model"
	"Food-delivery/pubsub"
	"context"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error)
}

// type IncreaseLikeCountStore interface {
// 	IncreaseLikeCount(ctx context.Context, id int) error
// }

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// increaseLikeStore IncreaseLikeCountStore
	ps pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		// increaseLikeStore: increaseLikeStore,
		ps: ps,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {

	_, err := biz.store.Find(ctx, map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId})
	// err := biz.store.Create(ctx, data)

	if err == nil {
		return restaurantlikemodel.ErrAlreadyLikedRestaurant()
	}

	err = biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// side messager
	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	//side effect

	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	// go func() {
	// 	defer common.AppRecover()
	// 	if err := biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// Side Effect

	// without Job
	//go func() {
	//	defer common.AppRecover()
	//	_ = biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//}()

	// with Job
	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//
	//	})
	//
	//	job.SetRetryDurations([]time.Duration{time.Second * 3})
	//
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	// New solution: use pubsub
	// biz.pubsub.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data))

	return nil
}
