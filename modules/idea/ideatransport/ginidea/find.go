package ginidea

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewstore"
)

func FindIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		viewStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewFindIdeaBiz(store, viewStore, appCtx.GetPubSub())

		result, err := biz.FindIdeaBiz(c.Request.Context(), ideaId, userId)
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

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		viewStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewFindIdeaBiz(store, viewStore, appCtx.GetPubSub())

		result, err := biz.FindIdeaBizForStaff(c.Request.Context(), ideaId, userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
