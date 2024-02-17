package middleware

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(ctx *gin.Context) {

	// return func(c *gin.Context) {
	// 	u := c.MustGet(common.CurrentUser).(common.Requester)
	// 	hasFound := false
	// 	for _, item := range allowRoles {
	// 		if u.GetRole() == item {
	// 			hasFound = true
	// 			break
	// 		}
	// 	}

	// 	if !hasFound {
	// 		panic(common.ErrNoPermission(errors.New("invalid role user ")))
	// 	}

	// 	c.Next()
	// }
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		for i := range allowRoles {
			if requester.GetRole() == allowRoles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(nil))
	}
}
