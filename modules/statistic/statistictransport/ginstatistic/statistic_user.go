package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/comment/commentstore"
	"web/modules/idealikeview/idealikeviewstore"
	"web/modules/statistic/statisticbiz"
	"web/modules/statistic/statisticmodel"
	"web/modules/user/userstore"
)

func ListStatisticUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter statisticmodel.StatisticUser
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		likeViewStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticUserBiz(likeViewStore, userStore, commentStore)

		result, err := biz.StatisticUserBiz(c.Request.Context(), &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
