package ginreply

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/replycomment/replybiz"
	"web/modules/replycomment/replystore"
)

func SoftDeleteReply(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		replyId, err := strconv.Atoi(c.Param("reply_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		replyStore := replystore.NewSQLStore(appCtx.GetDatabase())
		biz := replybiz.NewDeleteReplyBiz(replyStore, appCtx.GetPubSub())

		if err := biz.DeleteReplyBiz(c.Request.Context(), replyId, userId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}
