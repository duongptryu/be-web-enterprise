package ginuser

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/user/userbiz"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

func UpdateUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		userUpdateBiz := userbiz.NewUpdateUserBiz(userStore)

		if err := userUpdateBiz.UpdateUserBiz(c.Request.Context(), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
