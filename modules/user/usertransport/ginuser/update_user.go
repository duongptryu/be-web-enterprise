package ginuser

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/department/departmentstore"
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
		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		userUpdateBiz := userbiz.NewUpdateUserBiz(userStore, departmentStore)

		if err := userUpdateBiz.UpdateUserBiz(c.Request.Context(), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}

func UpdateUserSelf(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet(common.KeyUserHeader).(int)

		var data usermodel.UserUpdateSelf

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		userUpdateBiz := userbiz.NewUpdateUserBiz(userStore, departmentStore)

		if err := userUpdateBiz.UpdateUserSelfBiz(c.Request.Context(), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
