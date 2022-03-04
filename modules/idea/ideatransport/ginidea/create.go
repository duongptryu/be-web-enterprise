package ginidea

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

func CreateIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data ideamodel.IdeaCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewCreateIdeaBiz(store)

		if err := biz.CreateIdeaBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
