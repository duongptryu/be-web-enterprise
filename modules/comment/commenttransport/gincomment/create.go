package gincomment

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/comment/commentbiz"
	"web/modules/comment/commentmodel"
	"web/modules/comment/commentstore"
	"web/modules/idea/ideastore"
	"web/modules/notification/notificationstore"
	"web/modules/user/userstore"
)

func CreateComment(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data commentmodel.CommentCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		data.UserId = c.MustGet(common.KeyUserHeader).(int)

		commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())
		ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		notificationStore := notificationstore.NewSQLStore(appCtx.GetDatabase())

		biz := commentbiz.NewCreateCommentBiz(ideaStore, commentStore, acaYearStore, appCtx.GetPubSub(), appCtx.GetMailProvider(), userStore, notificationStore)

		if err := biz.CreateCommentBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
