package userstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/user/usermodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type UserStore interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*usermodel.User, error)
	ListUser(ctx context.Context,
		condition map[string]interface{},
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]usermodel.User, error)
	UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error
}
