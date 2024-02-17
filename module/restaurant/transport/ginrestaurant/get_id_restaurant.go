package ginrestaurant

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restaurantbiz "Food-delivery/module/restaurant/biz"
	restaurantmodel "Food-delivery/module/restaurant/model"
	restaurantrepository "Food-delivery/module/restaurant/repository"
	restaurantstorage "Food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		go func() {
			defer common.AppRecover()
			panic("aaaa")
		}()
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		uid, err := common.FromBase58(c.Param("id"))

		var data restaurantmodel.Restaurant

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// This is an error from Go standard lib, so we need to wrap it by common.ErrInvalidRequest
		// cuz this error is not normalized
		// var data restaurantmodel.Restaurant
		// if err != nil {
		// NOTICE:
		// we should just set `panic` in the transportation/controller layer
		// If we set `panic` in the business/services layer, because `panic`'s mechanism
		// will stop any code below it, so we might miss some logic in it
		// panic(common.ErrInvalidRequest(err))
		// }
		// data.UserId = requester.GetUserId()
		store := restaurantstorage.NewSQLStore(db)
		repo := restaurantrepository.NewGetRestaurantRepo(store, requester)
		biz := restaurantbiz.NewGetRestaurantBiz(repo)

		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			// Any err thrown from Biz belongs to Application error
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
		// c.HTML(http.StatusOK, "user.html", gin.H{
		// 	"Name": data.Name,
		// })
	}
}
