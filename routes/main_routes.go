package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	"Food-delivery/module/restaurant/transport/ginrestaurant"
	"Food-delivery/module/upload/transport/ginupload"
	"Food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.Upload(appContext))
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appContext))
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.GetRestaurant(appContext))
	restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))
	restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
}
