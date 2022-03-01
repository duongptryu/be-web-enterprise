package main

import (
	component "web/components"
	"web/middleware"
	"web/modules/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	r.Use(middleware.Recover(appCtx))
	v1Route(r, appCtx)
}

func v1Route(r *gin.Engine, appCtx component.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", ginuser.UserLogin(appCtx))
		user := v1.Group("/user", middleware.RequireAuthUserFeature(appCtx))
		{
			user.GET("", ginuser.ListUser(appCtx))
			user.POST("", ginuser.CreateUser(appCtx))
			user.PUT("/:user_id", ginuser.UpdateUser(appCtx))
			user.DELETE("/:user_id", ginuser.SoftDeleteUser(appCtx))
		}
		role := v1.Group("/role")
		{
			role.GET("", ginuser.GetListRole(appCtx))
		}
	}
}
