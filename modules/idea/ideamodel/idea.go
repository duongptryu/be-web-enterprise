package ideamodel

import "web/common"

const EntityName = "Idea"

type Idea struct {
	common.SQLModel
	Title         string        `json:"title" gorm:"column:title"`
	Content       string        `json:"content" gorm:"column:content"`
	UserId        int           `json:"user_id" gorm:"column:user_id"`
	CategoryId    int           `json:"category_id" gorm:"column:category_id"`
	AcaYearId     int           `json:"aca_year_id" gorm:"column:aca_year_id"`
	LikesCount    int           `json:"likes_count" gorm:"column:likes_count"`
	DislikesCount int           `json:"dislikes_count" gorm:"column:dislikes_count"`
	ViewsCount    int           `json:"views_count" gorm:"column:views_count"`
	CommentsCount int           `json:"comments_count" gorm:"column:comments_count"`
	Status        bool          `json:"status" gorm:"column:status"`
	Files         *common.Files `json:"files" gorm:"column:files"`
	IsAnonymous   bool          `json:"is_anonymous" gorm:"column:is_anonymous"`
	IsExpire      bool          `json:"is_expire" gorm:"is_expire"`
}

func (Idea) TableName() string {
	return "ideas"
}

type IdeaCreate struct {
	common.SQLModelCreate
	Title         string        `json:"title" gorm:"column:title"`
	Content       string        `json:"content" gorm:"column:content"`
	UserId        int           `json:"-" gorm:"column:user_id"`
	CategoryId    int           `json:"category_id" gorm:"column:category_id"`
	AcaYearId     int           `json:"-" gorm:"column:aca_year_id"`
	LikesCount    int           `json:"likes_count" gorm:"column:likes_count"`
	DislikesCount int           `json:"dislikes_count" gorm:"column:dislikes_count"`
	ViewsCount    int           `json:"views_count" gorm:"column:views_count"`
	CommentsCount int           `json:"comments_count" gorm:"column:comments_count"`
	Status        bool          `json:"status" gorm:"column:status"`
	Files         *common.Files `json:"files" gorm:"column:files"`
	IsAnonymous   bool          `json:"is_anonymous" gorm:"column:is_anonymous"`
	IsExpire      bool          `json:"-" gorm:"is_expire"`
}

func (IdeaCreate) TableName() string {
	return Idea{}.TableName()
}

type IdeaUpdate struct {
	common.SQLModelUpdate
	Title       string `json:"title" gorm:"column:title"`
	Content     string `json:"content" gorm:"column:content"`
	Status      bool   `json:"status" gorm:"column:status"`
	IsAnonymous bool   `json:"is_anonymous" gorm:"column:is_anonymous"`
	IsExpire    bool   `json:"-" gorm:"is_expire"`
}

func (IdeaUpdate) TableName() string {
	return Idea{}.TableName()
}
