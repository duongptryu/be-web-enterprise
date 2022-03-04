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

func (biz *updateIdeaBiz) UpdateIdeaBiz(ctx context.Context, acaYearId int, data *ideamodel.IdeaUpdate) error {
	cateDB, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": acaYearId})
	if err != nil {
		return common.ErrCannotUpdateEntity(ideamodel.EntityName, err)
	}
	if cateDB.Id == 0 {
		return common.ErrDataNotFound(ideamodel.EntityName)
	}

	//create user in db
	if err := biz.store.UpdateIdea(ctx, acaYearId, data); err != nil {
		return common.ErrCannotUpdateEntity(ideamodel.EntityName, err)
	}

	return nil
}