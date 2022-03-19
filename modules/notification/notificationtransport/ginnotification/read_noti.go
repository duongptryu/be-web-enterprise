package ginnotification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/common"
	component "web/components"
	"web/modules/notification/notificationbiz"
	"web/modules/notification/notificationstore"
)

func ReadNotification(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		notiId, err := strconv.Atoi(c.Param("noti_id"))
		if err != nil {
			panic(common.ErrParseJson(err))
		}

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := notificationstore.NewSQLStore(appCtx.GetDatabase())
		biz := notificationbiz.NewReadNotiBiz(store)

		if err := biz.ReadNotiBiz(c.Request.Context(), userId, notiId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(true))
	}
}
