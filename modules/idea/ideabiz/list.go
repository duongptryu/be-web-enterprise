package ideabiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type listIdeaBiz struct {
	store ideastore.IdeaStore
}

func NewListIdeaBiz(store ideastore.IdeaStore) *listIdeaBiz {
	return &listIdeaBiz{
		store: store,
	}
}

func (biz *listIdeaBiz) ListIdeaBiz(ctx context.Context, paging *common.Paging, filter *ideamodel.Filter) ([]ideamodel.Idea, error) {
	result, err := biz.store.ListIdea(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	return result, nil
}

func (biz *listIdeaBiz) ListIdeaBizForStaff(ctx context.Context, paging *common.Paging, filter *ideamodel.Filter) ([]ideamodel.Idea, error) {
	result, err := biz.store.ListIdea(ctx, map[string]interface{}{"status": true}, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	return result, nil
}
