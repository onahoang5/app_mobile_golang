package foodlikegin

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodlikebiz "Food-delivery/module/foodlike/biz"
	foodlikemodel "Food-delivery/module/foodlike/model"
	foodlikestorage "Food-delivery/module/foodlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserLikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := foodlikemodel.Filter{
			FoodId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := foodlikestorage.NewSQLStore(db)
		biz := foodlikebiz.NewListUserLikeFood(store)

		users, err := biz.ListUserLike(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
