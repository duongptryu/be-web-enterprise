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

func (biz *deleteIdeaBiz) DeleteIdeaBiz(ctx context.Context, cateId int) error {
	cateDB, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": cateId})
	if err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}
	if cateDB.Id == 0 {
		return common.ErrDataNotFound(ideamodel.EntityName)
	}

	//create user in db
	if err := biz.store.DeleteIdea(ctx, cateId); err != nil {
		return common.ErrCannotDeleteEntity(ideamodel.EntityName, err)
	}

	return nil
}
