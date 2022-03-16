package commentmodel

import "web/common"

const (
	EntityName = "Comment"
)

type Comment struct {
	common.SQLModel
	Id           int                `json:"id" gorm:"column:id"`
	UserId       int                `json:"user_id" gorm:"column:user_id"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
	IdeaId       int                `json:"idea_id" gorm:"column:idea_id"`
	Content      string             `json:"content" gorm:"column:content"`
	RepliesCount int                `json:"replies_count" gorm:"column:replies_count"`
	Status       bool               `json:"status" gorm:"column:status"`
	IsAnonymous  bool               `json:"is_anonymous" gorm:"column:is_anonymous"`
}

func (Comment) TableName() string {
	return "comments"
}

type CommentCreate struct {
	common.SQLModelCreate
	Id           int    `json:"-" gorm:"column:id"`
	UserId       int    `json:"-" gorm:"column:user_id"`
	IdeaId       int    `json:"idea_id" gorm:"column:idea_id"`
	Content      string `json:"content" gorm:"column:content" binding:"required"`
	RepliesCount int    `json:"-" gorm:"column:replies_count"`
	Status       bool   `json:"-" gorm:"status"`
	IsAnonymous  bool   `json:"is_anonymous" gorm:"column:is_anonymous"`
}

func (CommentCreate) TableName() string {
	return Comment{}.TableName()
}
