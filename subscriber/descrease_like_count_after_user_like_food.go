package subscriber

import (
	"Food-delivery/component/appctx"
	foodstorage "Food-delivery/module/food/storage"
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

func DescreaseLikeCountAfterUserLikeFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Des crease like count after user dislike restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMainDBConnection()
			store := foodstorage.NewSQLStore(db)
			likeData := message.Data().(HasFoodId)
			return store.DecreaseLikeCountFood(ctx, likeData.GetFoodId())
		},
	}
}
