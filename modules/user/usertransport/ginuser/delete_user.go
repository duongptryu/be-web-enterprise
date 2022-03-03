package ginuser

import (
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/user/userbiz"
	"web/modules/user/userstore"

	"github.com/gin-gonic/gin"
)

func SoftDeleteUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		userUpdateBiz := userbiz.NewDeleteUserBiz(userStore)

		if err := userUpdateBiz.SoftDeleteUserBiz(c.Request.Context(), userId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
