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

func LikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := foodlikemodel.Likefood{
			FoodId: int(uid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := foodlikestorage.NewSQLStore(db)
		// increstore := foodstorage.NewSQLStore(db)
		ps := appCtx.GetPubsub()
		biz := foodlikebiz.NewUserLikeFoodBiz(store, ps)

		if err = biz.LikeFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
