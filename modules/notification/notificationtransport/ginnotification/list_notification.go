package ginnotification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	component "web/components"
	"web/modules/notification/notificationbiz"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
)

func ListNotification(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter notificationmodel.Filter

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrParseJson(err))
		}
		paging.Fulfill()

		userId := c.MustGet(common.KeyUserHeader).(int)

		store := notificationstore.NewSQLStore(appCtx.GetDatabase())
		biz := notificationbiz.NewListNotificationBiz(store)

		result, err := biz.ListNotificationBiz(c.Request.Context(), userId, &paging, &filter)
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
