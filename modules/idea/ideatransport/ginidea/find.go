package ginidea

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideastore"
)

func FindIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewFindIdeaBiz(store)

		result, err := biz.FindIdeaBiz(c.Request.Context(), ideaId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}

func FindIdeaForStaff(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewFindIdeaBiz(store)

		result, err := biz.FindIdeaBizForStaff(c.Request.Context(), ideaId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
