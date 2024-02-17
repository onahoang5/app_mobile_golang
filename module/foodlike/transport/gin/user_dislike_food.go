package foodlikegin

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodlikebiz "Food-delivery/module/foodlike/biz"
	foodlikestorage "Food-delivery/module/foodlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisLikeFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		db := appCtx.GetMainDBConnection()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := foodlikestorage.NewSQLStore(db)
		// dercstore := foodstorage.NewSQLStore(db)
		ps := appCtx.GetPubsub()
		biz := foodlikebiz.NewDisLikeFoodBiz(store, ps)

		if err := biz.DisLikeFood(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
