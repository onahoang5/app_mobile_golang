package ginfood

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	foodbiz "Food-delivery/module/food/biz"
	foodmodel "Food-delivery/module/food/model"
	foodstorage "Food-delivery/module/food/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		// requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := foodmodel.FoodCreate{
			RestaurantId: int(uid.GetLocalID()),
		}
		if err != nil {
			log.Println(err)
		}

		if err := c.ShouldBind(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			panic(err)
		}

		store := foodstorage.NewSQLStore(db)
		biz := foodbiz.NewCreateFoodtBiz(store)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask(false)
		data.GenUID(common.DbTypeFood)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
