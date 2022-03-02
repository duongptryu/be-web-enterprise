package gincategory

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/category/categorybiz"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

func UpdateCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		cateId, err := strconv.Atoi(c.Param("cate_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		var data categorymodel.CategoryUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err := biz.UpdateCategoryBiz(c.Request.Context(), cateId, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}