package categorybiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

type createCategoryBiz struct {
	store categorystore.CategoryStore
}

func NewCreateCategoryBiz(store categorystore.CategoryStore) *createCategoryBiz {
	return &createCategoryBiz{
		store: store,
	}
}

func (biz *createCategoryBiz) CreateCategoryBiz(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
