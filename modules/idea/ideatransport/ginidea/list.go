package ginidea

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewstore"
)

func ListIdea(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter ideamodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userID := c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		likeStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewListIdeaBiz(store, likeStore)

		result, err := biz.ListIdeaBiz(c.Request.Context(), userID, &paging, &filter)
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

func ListIdeaForStaff(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter ideamodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userID := c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		likeStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewListIdeaBiz(store, likeStore)

		result, err := biz.ListIdeaBizForStaff(c.Request.Context(), userID, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListAllIdeaOwner(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter ideamodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		likeStore := idealikeviewstore.NewSQLStore(appCtx.GetDatabase())
		biz := ideabiz.NewListIdeaBiz(store, likeStore)

		result, err := biz.ListIdeaForOwner(c.Request.Context(), userId, &paging, &filter)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
