package idealikeviewbiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
	"web/modules/notification/notificationmodel"
	"web/modules/notification/notificationstore"
	"web/pubsub"
)

type createUserLikeViewIdeaBiz struct {
	store             idealikeviewstore.UserLikeViewIdeaStore
	ideaStore         ideastore.IdeaStore
	pubSub            pubsub.PubSub
	notificationStore notificationstore.NotificationStore
}

func NewCreateIdeaBiz(store idealikeviewstore.UserLikeViewIdeaStore, pubSub pubsub.PubSub, ideaStore ideastore.IdeaStore, notificationStore notificationstore.NotificationStore) *createUserLikeViewIdeaBiz {
	return &createUserLikeViewIdeaBiz{
		store:             store,
		pubSub:            pubSub,
		ideaStore:         ideaStore,
		notificationStore: notificationStore,
	}
}

func (biz *createUserLikeViewIdeaBiz) CreateUserLikeIdeaBiz(ctx context.Context, data *idealikeviewmodel.UserLikeIdea) error {
	existLike, err := biz.store.FindUserLikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId})
	if err != nil {
		return err
	}
	if existLike.IdeaId != 0 {
		return idealikeviewmodel.ErrUserAlreadyLikeIdea
	}

	existDislike, err := biz.store.FindUserDislikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId})
	if err != nil {
		return err
	}
	if existDislike.IdeaId != 0 {
		//remove dislike in database and Decrease dislike count
		if err := biz.store.DeleteUserDislikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId}); err != nil {
			return common.ErrCannotDeleteEntity(idealikeviewmodel.EntityName, err)
		}

		go biz.pubSub.Publish(ctx, common.TopicDecreaseDisLikeCountIdea, pubsub.NewMessage(data.IdeaId))
	}

	if err := biz.store.CreateUserLikeIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(idealikeviewmodel.EntityName, err)
	}

	go func() {
		go biz.pubSub.Publish(ctx, common.TopicIncreaseLikeCountIdea, pubsub.NewMessage(data.IdeaId))

		idea, _ := biz.ideaStore.FindIdea(ctx, map[string]interface{}{"id": data.IdeaId})

		if idea.UserId == data.UserId {
			return
		}

		newNoti := notificationmodel.NotificationIdeaCreate{
			TypeNoti: common.NewLikeIdeaNotification,
			OwnerId:  idea.UserId,
			IdeaId:   idea.Id,
			UserId:   data.UserId,
		}
		if err := biz.notificationStore.CreateNotification(ctx, &newNoti); err != nil {
			log.Error(err)
		}
	}()

	return nil
}

func (biz *createUserLikeViewIdeaBiz) CreateUserDislikeIdeaBiz(ctx context.Context, data *idealikeviewmodel.UserDislikeIdea) error {
	existDislike, err := biz.store.FindUserDislikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId})
	if err != nil {
		return err
	}
	if existDislike.IdeaId != 0 {
		return idealikeviewmodel.ErrUserAlreadyDisLikeIdea
	}

	existLike, err := biz.store.FindUserLikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId})
	if err != nil {
		return err
	}
	if existLike.IdeaId != 0 {
		//remove like in database and Decrease like count
		if err := biz.store.DeleteUserLikeIdea(ctx, map[string]interface{}{"user_id": data.UserId, "idea_id": data.IdeaId}); err != nil {
			return common.ErrCannotDeleteEntity(idealikeviewmodel.EntityName, err)
		}

		go biz.pubSub.Publish(ctx, common.TopicDecreaseLikeCountIdea, pubsub.NewMessage(data.IdeaId))
	}

	if err := biz.store.CreateUserDislikeIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(idealikeviewmodel.EntityName, err)
	}

	go func() {
		go biz.pubSub.Publish(ctx, common.TopicIncreaseDisLikeCountIdea, pubsub.NewMessage(data.IdeaId))

		idea, _ := biz.ideaStore.FindIdea(ctx, map[string]interface{}{"id": data.IdeaId})

		if idea.UserId == data.UserId {
			return
		}

		newNoti := notificationmodel.NotificationIdeaCreate{
			TypeNoti: common.NewDislikeNotification,
			OwnerId:  idea.UserId,
			IdeaId:   idea.Id,
			UserId:   data.UserId,
		}
		if err := biz.notificationStore.CreateNotification(ctx, &newNoti); err != nil {
			log.Error(err)
		}
	}()

	return nil
}
