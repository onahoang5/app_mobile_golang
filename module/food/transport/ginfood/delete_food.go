package ginfood

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodbiz "Food-delivery/module/food/biz"
	foodmodel "Food-delivery/module/food/model"
	foodstorage "Food-delivery/module/food/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		var data foodmodel.FoodUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewDeleteFoodBiz(store)

		if err := biz.DeleteFood(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
