package restautantlikebiz

import (
	"Food-delivery/common"
	restaurantlikemodel "Food-delivery/module/restaurantlike/model"
	"Food-delivery/pubsub"
	"context"
	"log"
)

type UserDisLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restautantId int) error
	Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error)
}

// type DecreaseLikeCountStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userDisLikeRestaurantBiz struct {
	store UserDisLikeRestaurantStore
	// decreaseLikeCountStore DecreaseLikeCountStore
	ps pubsub.Pubsub
}

func NewUserDisLikeRestaurantBiz(
	store UserDisLikeRestaurantStore,
	ps pubsub.Pubsub,
) *userDisLikeRestaurantBiz {
	return &userDisLikeRestaurantBiz{
		store: store,
		// decreaseLikeCountStore: decreaseLikeCountStore,
		ps: ps,
		// pubsub: pubsub,
	}
}

func (biz *userDisLikeRestaurantBiz) DisLikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {

	_, err := biz.store.Find(ctx, map[string]interface{}{"user_id": userId, "restaurant_id": restaurantId})

	if err != nil {
		return restaurantlikemodel.ErrAlreadyUnLikedRestaurant()
	}

	err = biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	//side effect
	if err := biz.ps.Publish(ctx, common.TopicUserDisLikeRestaurant,
		pubsub.NewMessage(&restaurantlikemodel.Like{
			RestaurantId: restaurantId,
			UserId:       userId,
		})); err != nil {
		log.Println(err)
	}

	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	// go func() {
	// 	defer common.AppRecover()
	// 	if err := biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
	// 		log.Println(err)

	// 		for i := 1; i <= 3; i++ {
	// 			err := biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	// 			if err == nil {
	// 				break
	// 			}
	// 			time.Sleep(time.Second * 3)
	// 		}
	// 	}
	// }()

	// Side Effect
	//without Job
	//go func() {
	//
	//	defer common.AppRecover()
	//	_ = biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	//}()

	//with Job
	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	//
	//	})
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	// Pubsub
	// biz.pubsub.Publish(
	// 	ctx,
	// 	common.TopicUserDislikeRestaurant,
	// 	pubsub.NewMessage(&restaurantlikemodel.Like{
	// 		RestaurantId: restaurantId,
	// 		UserId:       userId,
	// 	}))

	return nil
}
