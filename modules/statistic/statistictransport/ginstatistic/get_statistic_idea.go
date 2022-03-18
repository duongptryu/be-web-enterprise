package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticbiz"
	"web/modules/statistic/statisticmodel"
)

func ListStatisticIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter statisticmodel.StatisticReq
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticIdeaBiz(store)

		result, err := biz.StatisticIdeaBiz(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
