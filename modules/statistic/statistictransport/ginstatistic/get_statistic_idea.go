package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
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
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticIdeaBiz(store, acaYearStore)

		result, err := biz.StatisticIdeaBiz(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}

func ListStatisticIdeaByDay(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter statisticmodel.StatisticReq
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		v := &filter
		if v == nil {
			panic(common.ErrParseJson(common.NewCustomError(nil, "Academic year id not found", "")))
		}
		if v.AcaYearId == 0 {
			panic(common.ErrParseJson(common.NewCustomError(nil, "Academic year id not found", "")))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticIdeaBiz(store, acaYearStore)

		result, err := biz.StatisticIdeaByDayBiz(c.Request.Context(), filter.AcaYearId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
