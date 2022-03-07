package ginreply

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/comment/commentstore"
	"web/modules/replycomment/replybiz"
	"web/modules/replycomment/replymodel"
	"web/modules/replycomment/replystore"
)

func CreateReply(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data replymodel.ReplyCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		data.UserId = c.MustGet(common.KeyUserHeader).(int)

		commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())
		replyStore := replystore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := replybiz.NewCreateReplyBiz(replyStore, commentStore, acaYearStore, appCtx.GetPubSub())

		if err := biz.CreateReplyBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
