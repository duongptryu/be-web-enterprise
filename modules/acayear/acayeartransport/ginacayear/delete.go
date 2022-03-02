package ginacayear

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearbiz"
	"web/modules/acayear/acayearstore"
)

func DeleteAcaYear(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		acaYearId, err := strconv.Atoi(c.Param("aca_year_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}


		store := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := acayearbiz.NewDeleteAcaYearBiz(store)

		if err := biz.DeleteAcaYearBiz(c.Request.Context(), acaYearId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
