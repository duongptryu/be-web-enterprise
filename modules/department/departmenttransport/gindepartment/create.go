package gindepartment

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/department/departmentbiz"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
	"web/modules/user/userstore"
)

func CreateDepartment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data departmentmodel.DepartmentCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		userStore := userstore.NewSQLStore(appCtx.GetDatabase())

		biz := departmentbiz.NewCreateDepartmentBiz(departmentStore, userStore)

		if err := biz.CreateDepartment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
