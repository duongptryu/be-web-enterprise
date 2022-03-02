package categorystore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/category/categorymodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type CategoryStore interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
	DeleteCategory(ctx context.Context, id int) error
	FindCategory(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*categorymodel.Category, error)
	ListCategory(ctx context.Context,
		condition map[string]interface{},
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]categorymodel.Category, error)
	UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
}