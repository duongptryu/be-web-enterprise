package ginstatistic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/comment/commentstore"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewstore"
	"web/modules/statistic/statisticbiz"
	"web/modules/user/userstore"
)

func ListStatisticEachUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		likeViewStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())
		ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := statisticbiz.NewStatisticEachUserBiz(likeViewStore, userStore, commentStore, ideaStore)

		result, err := biz.StatisticEachUserBiz(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
