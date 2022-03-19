package notificationstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/notification/notificationmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{
		db: db,
	}
}

type NotificationStore interface {
	CreateNotification(ctx context.Context, data *notificationmodel.NotificationIdeaCreate) error
	ListNotification(ctx context.Context,
		condition map[string]interface{},
		filter *notificationmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]notificationmodel.NotificationIdea, error)
	UpdateNotification(ctx context.Context, condition map[string]interface{}, data *notificationmodel.NotificationIdeaUpdate) error
}
