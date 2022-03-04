package ideabiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type createIdeaBiz struct {
	store ideastore.IdeaStore
}

func NewCreateIdeaBiz(store ideastore.IdeaStore) *createIdeaBiz {
	return &createIdeaBiz{
		store: store,
	}
}

func (biz *createIdeaBiz) CreateIdeaBiz(ctx context.Context, data *ideamodel.IdeaCreate) error {
	if err := biz.store.CreateIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
