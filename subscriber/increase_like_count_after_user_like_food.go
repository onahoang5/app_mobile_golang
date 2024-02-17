package subscriber

import (
	"Food-delivery/component/appctx"
	foodstorage "Food-delivery/module/food/storage"
	"Food-delivery/pubsub"

	"context"
)

type HasFoodId interface {
	GetFoodId() int
	GetUserId() int
}

func IncreaseLikeCountAfterUserFood(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "In crease like count after user like food",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			db := appCtx.GetMainDBConnection()
			store := foodstorage.NewSQLStore(db)
			likeData := message.Data().(HasFoodId)
			return store.IncreaseLikeCountFood(ctx, likeData.GetFoodId())
		},
	}
}

// func PusnoticationWhenUserLikeFood(appCtx appctx.AppContext) consumerJob {
// 	return consumerJob{
// 		Title: "Pus notication when user like food",
// 		Hld: func(ctx context.Context, message *pubsub.Message) error {
// 			likeData := message.Data().(HasRestaurantId)
// 			log.Println("Push notication when user likes food", likeData.GetRestaurantId())
// 			return nil
// 		},
// 	}
// }

// func EmitRealtimeAfterUserLikeFood(appCtx appctx.AppContext) consumerJob {
// 	return consumerJob{
// 		Title: "Realtime emit after user like food",
// 		Hld: func(ctx context.Context, message *pubsub.Message) error {

// 			likeData := message.Data().(HasRestaurantId)
// 			appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
// 			return nil
// 		},
// 	}
// }
