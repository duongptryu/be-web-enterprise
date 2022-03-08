package ginuser

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/department/departmentstore"
	"web/modules/user/userbiz"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

func CreateUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		userRegisterStore := userstore.NewSQLStore(appCtx.GetDatabase())
		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		userRegisterBiz := userbiz.NewCreateUserBiz(userRegisterStore, departmentStore)

		if err := userRegisterBiz.CreateUserBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
