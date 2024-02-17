package ginroutes

import (
	"Food-delivery/component/appctx"
	"Food-delivery/middleware"
	"Food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SettupRoutesAdmin(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RequireAuth(appContext), middleware.RoleRequired(appContext, "admin"))
	{
		admin.GET("", ginuser.Profile(appContext))

	}

}
