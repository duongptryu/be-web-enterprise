package gindepartment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/department/departmentbiz"
	"web/modules/department/departmentmodel"
	"web/modules/department/departmentstore"
)

func ListDepartment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter departmentmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		biz := departmentbiz.NewListDepartment(departmentStore)

		result, err := biz.ListDepartmentBiz(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
