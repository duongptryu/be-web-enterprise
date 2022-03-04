package idealikeviewbiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idealikeview/idealikeviewstore"
)

type deleteUserLikeViewIdeaStore struct {
	store idealikeviewstore.UserLikeViewIdeaStore
}

func NewDeleteIdeaBiz(store idealikeviewstore.UserLikeViewIdeaStore) *deleteUserLikeViewIdeaStore {
	return &deleteUserLikeViewIdeaStore{
		store: store,
	}
}

func (biz *deleteUserLikeViewIdeaStore) DeleteUserLikeIdea(ctx context.Context, ideaId int, userId int) error {
	//create user in db
	if err := biz.store.DeleteUserLikeIdea(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaId}); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	return nil
}

func (biz *deleteUserLikeViewIdeaStore) DeleteUserDislikeIdea(ctx context.Context, ideaId int, userId int) error {
	//create user in db
	if err := biz.store.DeleteUserDislikeIdea(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaId}); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	return nil
}
