package ideamodel

import (
	"fmt"
	"web/common"
	"web/modules/category/categorymodel"
)

const EntityName = "Idea"

type Idea struct {
	common.SQLModel
	Title         string                   `json:"title" gorm:"column:title"`
	Content       string                   `json:"content" gorm:"column:content"`
	UserId        int                      `json:"user_id" gorm:"column:user_id"`
	User          *common.SimpleUser       `json:"user" gorm:"preload:false"`
	CategoryId    int                      `json:"category_id" gorm:"column:category_id"`
	Category      *categorymodel.Category  `json:"category" gorm:"preload:false"`
	AcaYearId     int                      `json:"aca_year_id" gorm:"column:aca_year_id"`
	LikesCount    int                      `json:"likes_count" gorm:"column:likes_count"`
	DislikesCount int                      `json:"dislikes_count" gorm:"column:dislikes_count"`
	ViewsCount    int                      `json:"views_count" gorm:"column:views_count"`
	CommentsCount int                      `json:"comments_count" gorm:"column:comments_count"`
	Status        bool                     `json:"status" gorm:"column:status"`
	Files         *common.Files            `json:"files" gorm:"column:files"`
	IsAnonymous   bool                     `json:"is_anonymous" gorm:"column:is_anonymous"`
	DepartmentId  int                      `json:"department_id" gorm:"column:department_id"`
	Department    *common.SimpleDepartment `json:"department,omitempty" gorm:"preload:false"`
	ThumbnailUrl  string                   `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	IsLike        bool                     `json:"is_like" gorm:"-"`
	IsDislike     bool                     `json:"is_dislike" gorm:"-"`
}

func (Idea) TableName() string {
	return "ideas"
}

func (data *Idea) SetTags() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v", data.Title, data.Content, data.User.FullName, data.Category.Name, data.Department.Name, data.Status)
}

func (data *Idea) GetIdeaId() int {
	return data.Id
}

type IdeaCreate struct {
	common.SQLModelCreate
	Title         string        `json:"title" gorm:"column:title" binding:"required"`
	Content       string        `json:"content" gorm:"column:content" binding:"required"`
	UserId        int           `json:"-" gorm:"column:user_id"`
	ThumbnailUrl  string        `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	CategoryId    int           `json:"category_id" gorm:"column:category_id" binding:"required"`
	AcaYearId     int           `json:"-" gorm:"column:aca_year_id"`
	LikesCount    int           `json:"-" gorm:"column:likes_count"`
	DislikesCount int           `json:"-" gorm:"column:dislikes_count"`
	ViewsCount    int           `json:"-" gorm:"column:views_count"`
	CommentsCount int           `json:"-" gorm:"column:comments_count"`
	Status        bool          `json:"-" gorm:"column:status"`
	Files         *common.Files `json:"files" gorm:"column:files"`
	IsAnonymous   bool          `json:"is_anonymous" gorm:"column:is_anonymous"`
	DepartmentId  int           `json:"-" gorm:"column:department_id"`
	Tags          string        `json:"-" gorm:"column:tags"`
}

func (IdeaCreate) TableName() string {
	return Idea{}.TableName()
}

func (data *IdeaCreate) SetTags() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v", data.Title, data.Content, data.Status)
}

type IdeaUpdate struct {
	common.SQLModelUpdate
	Title        string        `json:"title" gorm:"column:title"`
	Content      string        `json:"content" gorm:"column:content"`
	ThumbnailUrl string        `json:"thumbnail_url" gorm:"column:thumbnail_url"`
	Status       *bool         `json:"-" gorm:"column:status"`
	Files        *common.Files `json:"files" gorm:"column:files"`
	Tags         string        `json:"-" gorm:"column:tags"`
}

func (IdeaUpdate) TableName() string {
	return Idea{}.TableName()
}

var ErrAcademicYearNotReady = common.NewFullErrorResponse(409, nil, "Academic year not ready for now, please try again in later", "Academic year not ready for now, please try again in later", "ErrAcademicYearNotReady")
var ErrFirstClosureDateExpired = common.NewFullErrorResponse(409, nil, "Time to submit new idea is expire, please contact your admin to get support", "Time to submit new idea is expire, please contact your admin to get support", "ErrFirstClosureDateExpired")
var ErrFinalClosureDateExpired = common.NewFullErrorResponse(409, nil, "Time to submit new comment is expire, please contact your admin to get support", "Time to submit new comment is expire, please contact your admin to get support", "ErrFinalClosureDateExpired")
var ErrAccountCannotCreateIdea = common.NewFullErrorResponse(409, nil, "Your account is not staff, please contact your admin to get support", "Your account is not staff, please contact your admin to get support", "ErrFinalClosureDateExpired")
var ErrIdeaAlreadyRemoved = common.NewFullErrorResponse(409, nil, "This post idea already removed", "This post idea already removed", "ErrIdeaAlreadyRemoved")
