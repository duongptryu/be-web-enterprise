package notificationbiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
)

type listNotificationBiz struct {
	store notificationstore.NotificationStore
}

func NewListNotificationBiz(store notificationstore.NotificationStore) *listNotificationBiz {
	return &listNotificationBiz{
		store: store,
	}
}

func (biz *listNotificationBiz) ListNotificationBiz(ctx context.Context, userId int, paging *common.Paging, filter *notificationmodel.Filter) ([]notificationmodel.NotificationIdea, error) {
	result, err := biz.store.ListNotification(ctx, map[string]interface{}{"owner_id": userId}, filter, paging, "Idea", "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(notificationmodel.EntityName, err)
	}

	countNewest := 0
	for i, _ := range result {
		if !result[i].IsSee {
			countNewest++
			if i == 0 {
				go func(index int) {
					if err := biz.store.UpdateNotification(ctx, map[string]interface{}{"id": index}, &notificationmodel.NotificationIdeaUpdate{IsSee: true}); err != nil {
						log.Error(err)
					}
				}(result[i].Id)
			}
		}
		if result[i].IsSee {
			break
		}
	}
	filter.Newest = countNewest
	return result, nil
}
