package idealikeviewbiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
	"web/pubsub"
)

type createUserLikeViewIdeaBiz struct {
	store  idealikeviewstore.UserLikeViewIdeaStore
	pubSub pubsub.PubSub
}

func NewCreateIdeaBiz(store idealikeviewstore.UserLikeViewIdeaStore, pubSub pubsub.PubSub) *createUserLikeViewIdeaBiz {
	return &createUserLikeViewIdeaBiz{
		store:  store,
		pubSub: pubSub,
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
		return idealikeviewmodel.ErrUserAlreadyDisLikeIdea
	}

	if err := biz.store.CreateUserLikeIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicStaffLikeIdea, pubsub.NewMessage(data.IdeaId))

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
		return idealikeviewmodel.ErrUserAlreadyLikeIdea
	}

	if err := biz.store.CreateUserDislikeIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	go biz.pubSub.Publish(ctx, common.TopicStaffDislikeIdea, pubsub.NewMessage(data.IdeaId))

	return nil
}
