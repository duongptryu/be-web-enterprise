package notificationbiz

import (
	"context"
	"web/common"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
)

type readNotiBiz struct {
	store notificationstore.NotificationStore
}

func NewReadNotiBiz(store notificationstore.NotificationStore) *readNotiBiz {
	return &readNotiBiz{
		store: store,
	}
}

func (biz *readNotiBiz) ReadNotiBiz(ctx context.Context, userId int, notiId int) error {
	updateNoti := notificationmodel.NotificationIdeaUpdate{IsRead: true}

	if err := biz.store.UpdateNotification(ctx, map[string]interface{}{"id": notiId, "owner_id": userId}, &updateNoti); err != nil {
		return common.ErrCannotListEntity(notificationmodel.EntityName, err)
	}

	return nil
}
