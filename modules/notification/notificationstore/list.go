package notificationstore

import (
	"context"
	"web/common"
	"web/modules/notification/notificationmodel"
)

func (s *sqlStore) ListNotification(ctx context.Context,
	condition map[string]interface{},
	filter *notificationmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]notificationmodel.NotificationIdea, error) {
	var result []notificationmodel.NotificationIdea

	db := s.db

	db = db.Table(notificationmodel.NotificationIdea{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if paging.Total == 0 {
		return []notificationmodel.NotificationIdea{}, nil
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
