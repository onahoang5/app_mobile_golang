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

func DeleteRatingRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		var data restaurantratingmodel.UpdateRatingRestaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantratingstorage.NewSQLStore(db)

		biz := restaurantratingbiz.NewDeleteRatingRestaurantBiz(store)

		if err := biz.DeleteRatingRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {

			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": 1})

	}
}
