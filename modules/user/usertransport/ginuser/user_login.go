package ginuser

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/user/userbiz"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

func UserLogin(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := userstore.NewSQLStore(appCtx.GetDatabase())

		biz := userbiz.NewLoginBiz(store, appCtx.GetTokenProvider(), 60*60*24*30)
		account, err := biz.UserLogin(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(account))
	}
}
