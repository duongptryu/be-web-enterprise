package idealikeviewmodel

import "web/common"

const EntityName = "UserLikeIdea"

type UserLikeIdea struct {
	common.SQLModelCreateNoId
	UserId int                `json:"userId" gorm:"column:user_id"`
	User   *common.SimpleUser `json:"user" gorm:"preload:false"`
	IdeaId int                `json:"idea_id" gorm:"column:idea_id"`
}

func (UserLikeIdea) TableName() string {
	return "users_like_ideas"
}

type UserDislikeIdea struct {
	common.SQLModelCreateNoId
	UserId int                `json:"userId" gorm:"column:user_id"`
	User   *common.SimpleUser `json:"user" gorm:"preload:false"`
	IdeaId int                `json:"idea_id" gorm:"column:idea_id"`
}

func (UserDislikeIdea) TableName() string {
	return "users_dislike_ideas"
}

type UserViewIdea struct {
	common.SQLModelCreateNoId
	UserId int                `json:"userId" gorm:"column:user_id"`
	User   *common.SimpleUser `json:"user" gorm:"preload:false"`
	IdeaId int                `json:"idea_id" gorm:"column:idea_id"`
}

func (UserViewIdea) TableName() string {
	return "users_view_ideas"
}

var ErrUserAlreadyLikeIdea = common.NewCustomError(nil, "User already like idea", "ErrUserAlreadyLikeIdea")
var ErrUserAlreadyDisLikeIdea = common.NewCustomError(nil, "User already dislike idea", "ErrUserAlreadyDisLikeIdea")
