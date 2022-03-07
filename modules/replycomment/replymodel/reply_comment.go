package replymodel

import "web/common"

const (
	EntityName = "Reply_Comment"
)

type Reply struct {
	common.SQLModel
	UserId    int                `json:"user_id" gorm:"column:user_id"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
	IdeaId    int                `json:"idea_id" gorm:"column:idea_id"`
	CommentId int                `json:"comment_id" gorm:"column:comment_id"`
	Content   string             `json:"content" gorm:"column:content"`
	Status    bool               `json:"status" gorm:"status"`
}

func (Reply) TableName() string {
	return "replies"
}

type ReplyCreate struct {
	common.SQLModelCreate
	UserId    int    `json:"-" gorm:"column:user_id"`
	CommentId int    `json:"comment_id" gorm:"column:comment_id"`
	IdeaId    int    `json:"-" gorm:"column:idea_id"`
	Content   string `json:"content" gorm:"column:content"`
	Status    bool   `json:"-" gorm:"status"`
}

func (ReplyCreate) TableName() string {
	return Reply{}.TableName()
}
