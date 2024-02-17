package ginrestaurant

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restaurantbiz "Food-delivery/module/restaurant/biz"
	restaurantmodel "Food-delivery/module/restaurant/model"
	restaurantstorage "Food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		// id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		//ShouldBind là Gin mang toàn bộ request nó bind vào
		if err := c.ShouldBind(&data); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
			// panic(err)
		}

		db.Where("id= ?", int(uid.GetLocalID())).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
			// panic(err)
		}
		// chú ý dùng panic ở từng ngoài cùng thôi transport
		// c.JSON(http.StatusOK, gin.H {

		// })
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}

}
