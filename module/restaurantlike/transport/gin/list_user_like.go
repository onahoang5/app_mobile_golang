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

func ListUsersLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		//myArr := []string{}
		//
		//fmt.Println(myArr[0])

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restautantlikebiz.NewListUsersLikeRestaurantBiz(store)

		users, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
