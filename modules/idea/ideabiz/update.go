package ideabiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type updateIdeaBiz struct {
	store ideastore.IdeaStore
}

func NewUpdateIdeaBiz(store ideastore.IdeaStore) *updateIdeaBiz {
	return &updateIdeaBiz{
		store: store,
	}
}

func (biz *updateIdeaBiz) UpdateIdeaBiz(ctx context.Context, ideaId int, userId int, data *ideamodel.IdeaUpdate) error {
	ideaDB, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": ideaId})
	if err != nil {
		return common.ErrCannotUpdateEntity(ideamodel.EntityName, err)
	}
	if ideaDB.Id == 0 {
		return common.ErrDataNotFound(ideamodel.EntityName)
	}

	if ideaDB.UserId != userId {
		return common.ErrPermissionDenied
	}

	//create user in db
	if err := biz.store.UpdateIdea(ctx, ideaId, data); err != nil {
		return common.ErrCannotUpdateEntity(ideamodel.EntityName, err)
	}

	return nil
}