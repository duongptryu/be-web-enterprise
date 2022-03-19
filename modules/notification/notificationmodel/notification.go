package notificationmodel

import "web/common"

const EntityName = "Notification"

type NotificationIdea struct {
	common.SQLModel
	TypeNoti string             `json:"type_noti" gorm:"column:type_noti"`
	OwnerId  int                `json:"owner_id" gorm:"column:owner_id"`
	IdeaId   int                `json:"idea_id" gorm:"column:idea_id"`
	Idea     *common.SimpleIdea `json:"idea" gorm:"preload:false"`
	UserId   int                `json:"user_id" gorm:"column:user_id"`
	User     *common.SimpleUser `json:"user" gorm:"preload:false"`
	IsSee    bool               `json:"-" gorm:"column:is_see"`
	IsRead   bool               `json:"is_read" gorm:"column:is_read"`
}

func (NotificationIdea) TableName() string {
	return "notifications"
}

type NotificationIdeaCreate struct {
	common.SQLModelCreate
	TypeNoti string `json:"type_noti" gorm:"column:type_noti"`
	OwnerId  int    `json:"owner_id" gorm:"column:owner_id"`
	IdeaId   int    `json:"idea_id" gorm:"column:idea_id"`
	UserId   int    `json:"user_id" gorm:"column:user_id"`
	IsSea    bool   `json:"is_sea" gorm:"column:is_see"`
	IsRead   bool   `json:"is_read" gorm:"column:is_read"`
}

func (NotificationIdeaCreate) TableName() string {
	return NotificationIdea{}.TableName()
}

type NotificationIdeaUpdate struct {
	common.SQLModelUpdate
	IsSee  bool `gorm:"column:is_see" gorm:"column:is_see"`
	IsRead bool `gorm:"column:is_read" gorm:"column:is_read"`
}

func (NotificationIdeaUpdate) TableName() string {
	return NotificationIdea{}.TableName()
}
