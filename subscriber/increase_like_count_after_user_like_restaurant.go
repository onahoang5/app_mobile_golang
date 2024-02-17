package subscriber

import (
	"Food-delivery/component/appctx"
	restaurantstorage "Food-delivery/module/restaurant/storage"
	"Food-delivery/pubsub"
	"log"

	"context"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
}

// func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	db := appCtx.GetMainDBConnection()
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
// 	store := restaurantstorage.NewSQLStore(db)
// 	go func() {
// 		for {
// 			defer common.AppRecover()
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "In crease like count after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMainDBConnection()
			store := restaurantstorage.NewSQLStore(db)
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PusnoticationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Pus notication when user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notication when user likes restaurant", likeData.GetRestaurantId())
			return nil
		},
	}
}

func EmitRealtimeAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Realtime emit after user like restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {

			likeData := message.Data().(HasRestaurantId)
			appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
			return nil
		},
	}
}
