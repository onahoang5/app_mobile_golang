package ginfood

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodbiz "Food-delivery/module/food/biz"
	"Food-delivery/module/food/foodrepo"
	foodstorage "Food-delivery/module/food/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			defer common.AppRecover()
			panic("aaaa")
		}()

		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {

			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(db)
		repo := foodrepo.NewGetFoodRepo(store)
		biz := foodbiz.NewGetFoodBiz(repo)

		result, err := biz.GetFood(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			// Any err thrown from Biz belongs to Application error
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
