package ginrestaurantlike

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restautantlikebiz "Food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "Food-delivery/module/restaurantlike/model"
	restaurantlikestorage "Food-delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurants/:id/like

func LikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(db)
		// increaseLikeCountStore := restaurantstorage.NewSQLStore(db)
		//biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, increaseLikeCountStore, appCtx.GetPubsub())
		ps := appCtx.GetPubsub()
		biz := restautantlikebiz.NewUserLikeRestaurantBiz(store, ps)

		if err = biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)

		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
