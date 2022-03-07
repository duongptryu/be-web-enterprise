package idealikeviewbiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idealikeview/idealikeviewstore"
	"web/pubsub"
)

type deleteUserLikeViewIdeaStore struct {
	store  idealikeviewstore.UserLikeViewIdeaStore
	pubSub pubsub.PubSub
}

func NewDeleteIdeaBiz(store idealikeviewstore.UserLikeViewIdeaStore, pubSub pubsub.PubSub) *deleteUserLikeViewIdeaStore {
	return &deleteUserLikeViewIdeaStore{
		store:  store,
		pubSub: pubSub,
	}
}

func (biz *deleteUserLikeViewIdeaStore) DeleteUserLikeIdea(ctx context.Context, ideaId int, userId int) error {
	//create user in db
	if err := biz.store.DeleteUserLikeIdea(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaId}); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicDecreaseLikeCountIdea, pubsub.NewMessage(ideaId))

	return nil
}

func (biz *deleteUserLikeViewIdeaStore) DeleteUserDislikeIdea(ctx context.Context, ideaId int, userId int) error {
	//create user in db
	if err := biz.store.DeleteUserDislikeIdea(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaId}); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicDecreaseDisLikeCountIdea, pubsub.NewMessage(ideaId))

	return nil
}
