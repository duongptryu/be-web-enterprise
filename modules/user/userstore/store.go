package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"

	"gorm.io/gorm"
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
	SoftDeleteUser(ctx context.Context, id int) error
}
