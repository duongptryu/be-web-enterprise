package main

import (
	component "web/components"
	"web/middleware"
	"web/modules/acayear/acayeartransport/ginacayear"
	"web/modules/category/categorytransport/gincategory"
	"web/modules/comment/commenttransport/gincomment"
	"web/modules/department/departmenttransport/gindepartment"
	"web/modules/export/exporttransport/ginexport"
	"web/modules/idea/ideatransport/ginidea"
	"web/modules/idealikeview/idealikeviewtransport/ginuserlikeviewidea"
	"web/modules/replycomment/replytransport/ginreply"
	"web/modules/upload/uploadtransport/ginupload"
	"web/modules/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	r.Use(middleware.Recover(appCtx))
	r.Static("/assets", "/tmp")
	v1Route(r, appCtx)
}

func v1Route(r *gin.Engine, appCtx component.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", ginuser.UserLogin(appCtx))

		v1.GET("/export-csv", ginexport.ExportIdeaToCsv(appCtx))

		v1.GET("/export-docs", ginexport.ExportDocs(appCtx))

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

			department := admin.Group("/department")
			{
				department.GET("", gindepartment.ListDepartment(appCtx))
				department.POST("", gindepartment.CreateDepartment(appCtx))
				department.PUT("/:department_id", gindepartment.UpdateDepartment(appCtx))
			}
		}

		QAM := v1.Group("/qam", middleware.RequireQAMAuth(appCtx))
		{
			cate := QAM.Group("/category")
			{
				cate.GET("", gincategory.ListCategory(appCtx))
				cate.POST("", gincategory.CreateCategory(appCtx))
				cate.PUT("/:cate_id", gincategory.UpdateCategory(appCtx))
				cate.DELETE("/:cate_id", gincategory.DeleteCategory(appCtx))
			}
		}

		advance := v1.Group("adv", middleware.RequireAdvAuth(appCtx))
		{
			ideaAdv := advance.Group("/idea")
			{
				ideaAdv.GET("", ginidea.ListIdea(appCtx))
				ideaAdv.GET("/:idea_id", ginidea.FindIdea(appCtx))

				ideaAdv.GET("/:idea_id/comment", gincomment.ListCommentOfIdea(appCtx))

				ideaAdv.GET("/comment/:comment_id/reply", ginreply.ListReplyOfComment(appCtx))
			}
		}

		v1.Use(middleware.RequireAuth(appCtx))
		v1.POST("/upload", ginupload.Upload(appCtx))

		v1.GET("/profile", ginuser.GetProfileUser(appCtx))

		v1.GET("/category", gincategory.ListCategoryForStaff(appCtx))

		v1.GET("/department", gindepartment.ListDepartmentForStaff(appCtx))

		idea := v1.Group("/idea")
		{
			idea.GET("", ginidea.ListIdeaForStaff(appCtx))
			idea.GET("/:idea_id", ginidea.FindIdeaForStaff(appCtx))
			idea.GET("/my-idea", ginidea.ListAllIdeaOwner(appCtx))
			idea.POST("", ginidea.CreateIdea(appCtx))
			idea.PUT("/:idea_id", ginidea.UpdateIdea(appCtx))
			idea.DELETE("/:idea_id", ginidea.DeleteIdea(appCtx))

			//user like idea
			idea.POST("/:idea_id/like", ginuserlikeviewidea.CreateUserLikeIdea(appCtx))
			idea.DELETE("/:idea_id/like", ginuserlikeviewidea.DeleteUserLikeIdea(appCtx))
			//list user like idea
			idea.GET("/:idea_id/like", ginuserlikeviewidea.ListUserLikeIdea(appCtx))

			//user Dislike idea
			idea.POST("/:idea_id/dislike", ginuserlikeviewidea.CreateUserDislikeIdea(appCtx))
			idea.DELETE("/:idea_id/dislike", ginuserlikeviewidea.DeleteUserDislikeIdea(appCtx))
			//list user Dislike idea
			idea.GET("/:idea_id/dislike", ginuserlikeviewidea.ListUserDislikeIdea(appCtx))

			//list user view idea
			idea.GET("/:idea_id/view", ginuserlikeviewidea.ListUserViewIdea(appCtx))

			//comment
			idea.POST("/comment", gincomment.CreateComment(appCtx))
			idea.DELETE("/comment/:comment_id", gincomment.SoftDeleteComment(appCtx))
			idea.GET("/:idea_id/comment", gincomment.ListCommentOfIdeaForStaff(appCtx))

			//reply comment
			idea.POST("/comment/reply", ginreply.CreateReply(appCtx))
			idea.GET("/comment/:comment_id/reply", ginreply.ListReplyOfCommentForStaff(appCtx))
			idea.DELETE("/comment/reply/:reply_id", ginreply.SoftDeleteReply(appCtx))
		}
	}
}
