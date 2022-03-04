package ginuserlikeviewidea

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/idealikeview/idealikeviewbiz"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
)

func ListUserLikeIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter idealikeviewmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewListUserLikeViewIdeaBiz(store)

		result, err := biz.ListUserLikeIdea(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListUserDislikeIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter idealikeviewmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewListUserLikeViewIdeaBiz(store)

		result, err := biz.ListUserDislikeIdea(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListUserViewIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter idealikeviewmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := idealikeviewbiz.NewListUserLikeViewIdeaBiz(store)

		result, err := biz.ListUserViewIdea(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
