package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/user/userbiz"
	"web/modules/user/userstore"
)

func GetProfileUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdRaw := c.MustGet(common.KeyUserHeader)
		if err := userIdRaw.(int); err == 0 {
			panic(common.ErrUnAuthorization)
		}
		userIdInt := userIdRaw.(int)

		store := userstore.NewSQLStore(appCtx.GetDatabase())
		biz := userbiz.NewFindUserBiz(store)

		result, err := biz.FindUserBiz(c.Request.Context(), userIdInt)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(result))
	}
}
