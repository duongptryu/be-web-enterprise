package ginuserlikeviewidea

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/idealikeview/idealikeviewbiz"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
)

func CreateUserLikeIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data idealikeviewmodel.UserLikeIdea

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

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
		var data idealikeviewmodel.UserDislikeIdea

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewCreateIdeaBiz(store, appCtx.GetPubSub())

		if err := biz.CreateUserDislikeIdeaBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
