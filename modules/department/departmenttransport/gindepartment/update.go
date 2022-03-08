package gindepartment

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/department/departmentbiz"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
)

func UpdateDepartment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		departmentId, err := strconv.Atoi(c.Param("department_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data departmentmodel.DepartmentUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())

		biz := departmentbiz.NewUpdateIdeaBiz(departmentStore)

		if err := biz.UpdateIdeaBiz(c.Request.Context(), departmentId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}