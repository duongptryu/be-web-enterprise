package gincategory

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/category/categorybiz"
	"web/modules/category/categorystore"
)

func DeleteCategory(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		cateId, err := strconv.Atoi(c.Param("cate_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetDatabase())
		biz := categorybiz.NewDeleteCategoryBiz(store)

		if err := biz.DeleteCategoryBiz(c.Request.Context(), cateId); err != nil {
			panic(err)
		}

		c.JSON(200, common.NewSimpleSuccessResponse(true))
	}
}