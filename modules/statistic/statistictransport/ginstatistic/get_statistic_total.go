package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticbiz"
	"web/modules/user/userstore"
)

func ListStatisticTotal(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticTotalBiz(store, userStore)

		result, err := biz.StatisticTotalBiz(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
