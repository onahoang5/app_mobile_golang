package restaurantratinggin

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restaurantratingbiz "Food-delivery/module/restaurantrating/biz"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	restaurantratingstorage "Food-delivery/module/restaurantrating/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RatingRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		db := appCtx.GetMainDBConnection()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantratingmodel.CreateRatingRestaurant{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		if err := c.ShouldBind(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			panic(err)
		}

		store := restaurantratingstorage.NewSQLStore(db)
		biz := restaurantratingbiz.NewUserRatingRestaurantBiz(store)

		if err := biz.RatingRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeRestaurant)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
