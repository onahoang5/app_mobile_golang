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

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		// go func() {
		// 	// defer common.AppRecover()

		// 	arr := []int{}
		// 	log.Println(arr[0])
		// }()

		var data restaurantmodel.RestaurantCreate
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

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(err)
		}

		data.Mask(false)
		// chú ý dùng panic ở từng ngoài cùng thôi transport
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
