package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	foodlikegin "Food-delivery/module/foodlike/transport/gin"

	"github.com/gin-gonic/gin"
)

func SetupRoutesLikeFood(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	foods := v1.Group("/foods", middleware.RequireAuth(appContext))
	foods.POST("/:id/like", foodlikegin.LikeFood(appContext))
	foods.DELETE("/:id/dislike", foodlikegin.DisLikeFood(appContext))
	foods.GET("/:id/liked-users", foodlikegin.ListUserLikeFood(appContext))
}
