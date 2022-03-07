package ginreply

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/replycomment/replybiz"
	"web/modules/replycomment/replymodel"
	"web/modules/replycomment/replystore"
)

func ListReplyOfComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter replymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := replystore.NewSQLStore(appCtx.GetDatabase())
		biz := replybiz.NewListReplyBiz(store)

		result, err := biz.ListReplyBiz(c.Request.Context(), commentId, &paging, &filter)
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

func ListReplyOfCommentForStaff(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var filter replymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := replystore.NewSQLStore(appCtx.GetDatabase())
		biz := replybiz.NewListReplyBiz(store)

		result, err := biz.ListReplyForStaff(c.Request.Context(), commentId, &paging, &filter)
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
