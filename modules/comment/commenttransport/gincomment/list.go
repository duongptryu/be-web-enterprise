package gincomment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/comment/commentbiz"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
)

func ListCommentOfIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter commentmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := commentstore.NewSQLStore(appCtx.GetDatabase())
		biz := commentbiz.NewListComment(store)

		result, err := biz.ListComment(c.Request.Context(), ideaId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListCommentOfIdeaForStaff(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter commentmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := commentstore.NewSQLStore(appCtx.GetDatabase())
		biz := commentbiz.NewListComment(store)

		result, err := biz.ListCommentForStaff(c.Request.Context(), ideaId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		for i := range result {
			if i == len(result)-1 {
				paging.NextCursor = result[i].Id
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}