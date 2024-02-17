package subscriber

import (
	"Food-delivery/component/appctx"
	restaurantstorage "Food-delivery/module/restaurant/storage"
	"Food-delivery/pubsub"

	"context"
)

// func DescreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	db := appCtx.GetMainDBConnection()
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserDisLikeRestaurant)
// 	store := restaurantstorage.NewSQLStore(db)
// 	go func() {
// 		for {
// 			defer common.AppRecover()
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func DescreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Des crease like count after user dislike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMainDBConnection()
			store := restaurantstorage.NewSQLStore(db)
			likeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
