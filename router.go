package main

import (
	component "web/components"
	"web/middleware"
	"web/modules/acayear/acayeartransport/ginacayear"
	"web/modules/category/categorytransport/gincategory"
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

		role := v1.Group("/role")
		{
			role.GET("", ginuser.GetListRole(appCtx))
		}

		admin := v1.Group("/admin", middleware.RequireAdminAuth(appCtx))
		{
			academicYear := admin.Group("/academic-year")
			{
				academicYear.GET("", ginacayear.ListAcaYear(appCtx))
				academicYear.POST("", ginacayear.CreateAcademicYear(appCtx))
				academicYear.PUT("/:aca_year_id", ginacayear.UpdateAcademicYear(appCtx))
				academicYear.DELETE("/:aca_year_id", ginacayear.DeleteAcaYear(appCtx))
			}

			user := admin.Group("/user")
			{
				user.GET("", ginuser.ListUser(appCtx))
				user.POST("", ginuser.CreateUser(appCtx))
				user.PUT("/:user_id", ginuser.UpdateUser(appCtx))
				user.DELETE("/:user_id", ginuser.SoftDeleteUser(appCtx))
			}
		}

		QAM := v1.Group("/qam", middleware.RequireQAMAuth(appCtx))
		{
			cate := QAM.Group("/category")
			{
				cate.GET("", gincategory.ListCategory(appCtx))
				cate.POST("", gincategory.CreateCategory(appCtx))
				cate.PUT("/:cate_id", gincategory.UpdateCategory(appCtx))
				cate.DELETE("/:user_id", gincategory.DeleteCategory(appCtx))
			}
		}

		v1.Use(middleware.RequireAuth(appCtx))
		v1.GET("/profile", ginuser.GetProfileUser(appCtx))
		v1.GET("/category", gincategory.ListCategoryForStaff(appCtx))
	}
}
