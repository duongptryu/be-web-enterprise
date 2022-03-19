package notificationstore

import (
	"context"
	"web/common"
	"web/modules/notification/notificationmodel"
)

func (s *sqlStore) UpdateNotification(ctx context.Context, condition map[string]interface{}, data *notificationmodel.NotificationIdeaUpdate) error {
	db := s.db

	if err := db.Table(notificationmodel.NotificationIdeaUpdate{}.TableName()).Where(condition).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
