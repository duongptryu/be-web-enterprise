package idealikeviewstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type UserLikeViewIdeaStore interface {
	CreateUserLikeIdea(ctx context.Context, data *idealikeviewmodel.UserLikeIdea) error
	CreateUserDislikeIdea(ctx context.Context, data *idealikeviewmodel.UserDislikeIdea) error
	CreateUserViewIdea(ctx context.Context, data *idealikeviewmodel.UserViewIdea) error
	DeleteUserLikeIdea(ctx context.Context, condition map[string]interface{}) error
	DeleteUserDislikeIdea(ctx context.Context, condition map[string]interface{}) error
	ListUserLikeIdea(ctx context.Context,
		condition map[string]interface{},
		filter *idealikeviewmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]idealikeviewmodel.UserLikeIdea, error)
	ListUserDislikeIdea(ctx context.Context,
		condition map[string]interface{},
		filter *idealikeviewmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]idealikeviewmodel.UserDislikeIdea, error)
	ListUserViewIdea(ctx context.Context,
		condition map[string]interface{},
		filter *idealikeviewmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]idealikeviewmodel.UserViewIdea, error)
	FindUserViewIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*idealikeviewmodel.UserViewIdea, error)
	FindUserLikeIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*idealikeviewmodel.UserLikeIdea, error)
	FindUserDislikeIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*idealikeviewmodel.UserDislikeIdea, error)

	ListIdeaUserLike(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
	ListIdeaUserDislike(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
	CountUserLike(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
	CountUserDisLike(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
	CountUserViewIdea(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (map[int]int, error)
}
