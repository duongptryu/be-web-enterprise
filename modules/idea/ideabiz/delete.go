package ideabiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type deleteIdeaBiz struct {
	store ideastore.IdeaStore
}

func NewDeleteIdeaBiz(store ideastore.IdeaStore) *deleteIdeaBiz {
	return &deleteIdeaBiz{
		store: store,
	}
}

func (biz *deleteIdeaBiz) DeleteIdeaBiz(ctx context.Context, ideaId int, userId int) error {
	ideaDB, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": ideaId, "status": true})
	if err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}
	if ideaDB.Id == 0 {
		return common.ErrDataNotFound(ideamodel.EntityName)
	}

	if ideaDB.UserId != userId {
		return common.ErrPermissionDenied
	}

	//create user in db
	if err := biz.store.DeleteIdea(ctx, ideaId); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	return nil
}
