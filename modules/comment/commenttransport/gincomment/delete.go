package gincomment

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/comment/commentbiz"
	"web/modules/comment/commentstore"
)

func SoftDeleteComment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		commentId, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())
		biz := commentbiz.NewDeleteCommentBiz(commentStore, appCtx.GetPubSub())

		if err := biz.DeleteCommentBiz(c.Request.Context(), commentId, userId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
