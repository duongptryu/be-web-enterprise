package ginidea

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
	"web/modules/acayear/acayearstore"
	"web/modules/category/categorystore"
	"web/modules/department/departmentstore"
	"web/modules/idea/ideabiz"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/notification/notificationstore"
	"web/modules/user/userstore"
)

func CreateIdea(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data ideamodel.IdeaCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrParseJson(err))
		}

		data.UserId = c.MustGet(common.KeyUserHeader).(int)

		store := ideastore.NewSQLStore(appCtx.GetDatabase())
		categoryStore := categorystore.NewSQLStore(appCtx.GetDatabase())
		userStore := userstore.NewSQLStore(appCtx.GetDatabase())
		acaYearStore := acayearstore.NewSQLStore(appCtx.GetDatabase())
		departmentStore := departmentstore.NewSQLStore(appCtx.GetDatabase())
		notificationStore := notificationstore.NewSQLStore(appCtx.GetDatabase())

		biz := ideabiz.NewCreateIdeaBiz(store, categoryStore, acaYearStore, userStore, departmentStore, appCtx.GetMailProvider(), notificationStore)

		if err := biz.CreateIdeaBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(201, common.NewSimpleSuccessResponse(true))
	}
}
