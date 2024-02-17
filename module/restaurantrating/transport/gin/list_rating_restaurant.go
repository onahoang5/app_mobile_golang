package restaurantratinggin

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	restaurantratingbiz "Food-delivery/module/restaurantrating/biz"
	restaurantratingmodel "Food-delivery/module/restaurantrating/model"
	restaurantratingrepo "Food-delivery/module/restaurantrating/repo"
	restaurantratingstorage "Food-delivery/module/restaurantrating/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDataUserRatingRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		// uid, err := common.FromBase58(c.Param("id"))

		// if err != nil {
		// 	panic(common.ErrInvalidRequest(err))
		// }
		// filter := restaurantratingmodel.Filter{
		// 	RestaurantId: int(uid.GetLocalID()),
		// }

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		var filter restaurantratingmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {

			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}

		store := restaurantratingstorage.NewSQLStore(db)
		repo := restaurantratingrepo.NewListRatingRestaurantRepo(store)
		biz := restaurantratingbiz.NewListRatingRestaurantBiz(repo)

		users, err := biz.ListRatingRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
