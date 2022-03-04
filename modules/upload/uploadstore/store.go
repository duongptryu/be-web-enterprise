package uploadstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type FileStore interface {
	CreateFile(ctx context.Context, data *common.File) error
	DeleteFile(ctx context.Context, id int) error
	FindFile(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*common.File, error)
}
