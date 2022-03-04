package ideabiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type findIdeaBiz struct {
	store ideastore.IdeaStore
}

func NewFindIdeaBiz(store ideastore.IdeaStore) *findIdeaBiz {
	return &findIdeaBiz{
		store: store,
	}
}

func (biz *findIdeaBiz) FindIdeaBiz(ctx context.Context, id int) (*ideamodel.Idea, error) {
	result, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	return result, nil
}

func (biz *findIdeaBiz) FindIdeaBizForStaff(ctx context.Context, id int) (*ideamodel.Idea, error) {
	result, err := biz.store.FindIdea(ctx, map[string]interface{}{"id": id, "status": true})
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	return result, nil
}