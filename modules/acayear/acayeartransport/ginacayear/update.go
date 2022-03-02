package ginacayear

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearbiz"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

func UpdateAcademicYear(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		acaYearId, err := strconv.Atoi(c.Param("aca_year_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data acayearmodel.AcademicYearUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := acayearbiz.NewUpdateUserBiz(store)

		if err := biz.UpdateUserBiz(c.Request.Context(), acaYearId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
