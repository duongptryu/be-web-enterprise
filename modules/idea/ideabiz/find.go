package ideabiz

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewmodel"
	"web/modules/idealikeview/idealikeviewstore"
	"web/pubsub"
)

type findIdeaBiz struct {
	store     ideastore.IdeaStore
	viewStore idealikeviewstore.UserLikeViewIdeaStore
	pubSub    pubsub.PubSub
}

func NewFindIdeaBiz(store ideastore.IdeaStore, viewStore idealikeviewstore.UserLikeViewIdeaStore, pubSub pubsub.PubSub) *findIdeaBiz {
	return &findIdeaBiz{
		store:     store,
		viewStore: viewStore,
		pubSub:    pubSub,
	}
}

func (biz *findIdeaBiz) FindIdeaBiz(ctx context.Context, id int, userId int) (*ideamodel.Idea, error) {
	result, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": id}, "User", "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}
	if result.Id == 0 {
		return nil, common.ErrDataNotFound(ideamodel.EntityName)
	}

	//increase view count
	go biz.increaseViewCountIdea(ctx, id, userId)

	return result, nil
}

func (biz *findIdeaBiz) FindIdeaBizForStaff(ctx context.Context, id int, userId int) (*ideamodel.Idea, error) {
	result, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": id, "status": true}, "User", "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}
	if result.Id == 0 {
		return nil, common.ErrDataNotFound(ideamodel.EntityName)
	}

	//increase view count
	go biz.increaseViewCountIdea(ctx, id, userId)

	return result, nil
}

func (biz *findIdeaBiz) increaseViewCountIdea(ctx context.Context, ideaId, userId int) {
	exist, err := biz.viewStore.FindUserViewIdea(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaId})
	if err != nil {
		log.Error(err)
		return
	}
	if exist.IdeaId == 0 {
		err := biz.viewStore.CreateUserViewIdea(ctx, &idealikeviewmodel.UserViewIdea{IdeaId: ideaId, UserId: userId})
		if err != nil {
			log.Println(err)
		}
		go biz.pubSub.Publish(ctx, common.TopicIncreaseViewCountIdea, pubsub.NewMessage(ideaId))
	}
}
