package ginuser

import (
	"github.com/gin-gonic/gin"
	"web/common"
	component "web/components"
)

func GetListRole(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		listRole := []string{common.RoleAdmin, common.RoleQAManager, common.RoleQACoordinator, common.RoleStaff}

		c.JSON(200, common.NewSimpleSuccessResponse(listRole))
	}
}
