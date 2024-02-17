package ginuser

import (
	"Food-delivery/common"
	"Food-delivery/component/appctx"
	"Food-delivery/component/hasher"
	"Food-delivery/component/tokenprovider/jwt"
	userbiz "Food-delivery/module/user/biz"
	usermodel "Food-delivery/module/user/model"
	userstorage "Food-delivery/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey()) //appctx.SecretKey()
		// tokenProvider := jwt.NewTokenJWTProvider("CristianoRonaldo")
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(appCtx, store, 60*60*24*30, tokenProvider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
