package gincategory

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/category/categorybiz"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

func CreateCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategoryBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
