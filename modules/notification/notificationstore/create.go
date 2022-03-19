package notificationstore

import (
	"context"
	"web/common"
	"web/modules/notification/notificationmodel"
)

func (s *sqlStore) CreateNotification(ctx context.Context, data *notificationmodel.NotificationIdeaCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
