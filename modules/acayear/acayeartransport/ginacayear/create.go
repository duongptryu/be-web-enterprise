package ginacayear

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearbiz"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

func CreateAcademicYear(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data acayearmodel.AcademicYearCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := acayearbiz.NewCreateAcaYearBiz(store)

		if err := biz.CreateAcaYearBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
