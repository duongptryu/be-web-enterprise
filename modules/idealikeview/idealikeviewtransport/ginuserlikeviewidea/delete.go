package ginuserlikeviewidea

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/idealikeview/idealikeviewbiz"
	"web/modules/idealikeview/idealikeviewstore"
)

func DeleteUserLikeIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewDeleteIdeaBiz(store, appCtx.GetPubSub())

		if err := biz.DeleteUserLikeIdea(c.Request.Context(), ideaId, userIdRaw.(int)); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}

func DeleteUserDislikeIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		ideaId, err := strconv.Atoi(c.Param("idea_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userIdRaw := c.MustGet(common.KeyUserHeader)

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewDeleteIdeaBiz(store, appCtx.GetPubSub())

		if err := biz.DeleteUserDislikeIdea(c.Request.Context(), ideaId, userIdRaw.(int)); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
