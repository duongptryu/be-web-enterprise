package ginacayear

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearbiz"
	"web/modules/acayear/acayearstore"
)

func FindCurrentAcaYear(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := acayearstore.NewSQLStore(appCtx.GetDatabase())
		biz := acayearbiz.NewFindAcaYear(store)

		result, err := biz.FindAcaYear(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}