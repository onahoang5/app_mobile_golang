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

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fulfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		// likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepository.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": err.Error(),
			// })
			// return
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
