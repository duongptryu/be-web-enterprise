package ginuserlikeviewidea

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idealikeview/idealikeviewbiz"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
)

func CreateUserLikeIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data idealikeviewmodel.UserLikeIdea

		data.IdeaId = ideaId
		data.UserId = c.MustGet(common.KeyUserHeader).(int)

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewCreateIdeaBiz(store, appCtx.GetPubSub())

		if err := biz.CreateUserLikeIdeaBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}

func CreateUserDislikeIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data idealikeviewmodel.UserDislikeIdea

		data.IdeaId = ideaId
		data.UserId = c.MustGet(common.KeyUserHeader).(int)

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewCreateIdeaBiz(store, appCtx.GetPubSub())

		if err := biz.CreateUserDislikeIdeaBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
