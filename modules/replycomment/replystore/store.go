package replystore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/replycomment/replymodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type ReplyStore interface {
	SoftDeleteReply(ctx context.Context, id int) error
	DeleteReply(ctx context.Context, id int) error
	FindReply(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*replymodel.Reply, error)
	ListReply(ctx context.Context,
		condition map[string]interface{},
		filter *replymodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]replymodel.Reply, error)
	CreateReply(ctx context.Context, data *replymodel.ReplyCreate) error
}
