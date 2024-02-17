package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	restaurantratinggin "Food-delivery/module/restaurantrating/transport/gin"

	"github.com/gin-gonic/gin"
)

func SetupRoutesRatingRestaurant(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appContext))
	restaurants.POST("/:id/rating", restaurantratinggin.RatingRestaurant(appContext))
	restaurants.DELETE("/:id/comment", restaurantratinggin.DeleteRatingRestaurant(appContext))
	restaurants.GET("/listrestaurant", restaurantratinggin.ListDataUserRatingRestaurant(appContext))
	restaurants.GET("/:id/listid", restaurantratinggin.GetRatingRestaurant(appContext))
}
