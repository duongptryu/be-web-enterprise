package ginidea

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideastore"
)

func DeleteIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewDeleteIdeaBiz(store)

		if err := biz.DeleteIdeaBiz(c.Request.Context(), ideaId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
