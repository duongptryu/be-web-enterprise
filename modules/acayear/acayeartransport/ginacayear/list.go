package ginacayear

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearbiz"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

func ListAcaYear(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter acayearmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		query := c.Query("query")
		if err := paging.ParsePaging(query); err != nil {
			panic(err)
		}
		paging.Fulfill()

		store := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := acayearbiz.NewListAcaYearBiz(store)

		result, err := biz.ListAcaYearBiz(c.Request.Context(), &paging, &filter)
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
