package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/category/categorystore"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticbiz"
)

func ListCountIdeaByCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		categoryStore := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticCountIdeaInCategoryBiz(store, categoryStore)

		result, err := biz.StatisticCountIdeaInCategoryBiz(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
