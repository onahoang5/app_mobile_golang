package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	ginrestaurantlike "Food-delivery/module/restaurantlike/transport/gin"

	"github.com/gin-gonic/gin"
)

func SetupRoutesLike(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appContext))
	restaurants.POST("/:id/like", ginrestaurantlike.LikeRestaurant(appContext))
	restaurants.DELETE("/:id/dislike", ginrestaurantlike.DisLikeRestaurant(appContext))
	restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appContext))

}
