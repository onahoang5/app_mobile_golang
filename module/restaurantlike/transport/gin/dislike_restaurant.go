package ginrestaurantlike

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restautantlikebiz "Food-delivery/module/restaurantlike/biz"
	restaurantlikestorage "Food-delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurants/:id/unlike

func DisLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		db := appCtx.GetMainDBConnection()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(db)
		//decreaseLikeCountRestaurant := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		// decreaselike := restaurantstorage.NewSQLStore(db)
		// pubsub := appCtx.GetPubsub()
		ps := appCtx.GetPubsub()
		//biz := restaurantlikebiz.NewUserUnLikeRestaurantBiz(store, decreaseLikeCountRestaurant)
		biz := restautantlikebiz.NewUserDisLikeRestaurantBiz(store, ps)

		if err = biz.DisLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
