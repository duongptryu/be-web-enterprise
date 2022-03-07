package commentstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/comment/commentmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type CommentStore interface {
	CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error
	SoftDeleteComment(ctx context.Context, id int) error
	DeleteComment(ctx context.Context, id int) error
	ListComment(ctx context.Context,
		condition map[string]interface{},
		filter *commentmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]commentmodel.Comment, error)
	FindComment(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*commentmodel.Comment, error)
}
