package ginfood

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodbiz "Food-delivery/module/food/biz"
	"Food-delivery/module/food/foodrepo"
	foodmodel "Food-delivery/module/food/model"
	foodstorage "Food-delivery/module/food/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter foodmodel.Filter

		db := appCtx.GetMainDBConnection()

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{"ok": 1})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{"oke": 1})
			return

		}
		paging.Fulfill()

		store := foodstorage.NewSQLStore(db)
		repo := foodrepo.NewListFoodRepo(store)
		biz := foodbiz.NewListFoodtBiz(repo)

		result, err := biz.ListFood(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
