package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	"Food-delivery/module/food/transport/ginfood"

	"github.com/gin-gonic/gin"
)

func SetupRoutesFood(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	food := v1.Group("/foods", middleware.RequireAuth(appContext))
	food.POST("/:id", ginfood.CreateFood(appContext))
	food.DELETE("/:id/foods", ginfood.DeleteFood(appContext))
	food.GET("/", ginfood.ListFood(appContext))
	food.PATCH("/:id", ginfood.UpdateFood(appContext))
	food.GET("/:id", ginfood.GetFood(appContext))

}
