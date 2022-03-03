package gincategory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/category/categorybiz"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

func ListCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter categorymodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		store := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := categorybiz.NewListCategoryBiz(store)

		result, err := biz.ListCategoryBiz(c.Request.Context(), &paging, &filter)
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

func ListCategoryForStaff(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter categorymodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrParseJson(err))
		}

		var paging common.Paging
		query := c.Query("query")
		if err := paging.ParsePaging(query); err != nil {
			panic(err)
		}
		paging.Fulfill()

		store := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := categorybiz.NewListCategoryBiz(store)

		result, err := biz.ListCategoryBizForStaff(c.Request.Context(), &paging, &filter)
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

