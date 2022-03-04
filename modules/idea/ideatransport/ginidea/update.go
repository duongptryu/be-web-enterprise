package ginidea

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

func UpdateIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data ideamodel.IdeaUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewUpdateIdeaBiz(store)

		if err := biz.UpdateIdeaBiz(c.Request.Context(), ideaId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
