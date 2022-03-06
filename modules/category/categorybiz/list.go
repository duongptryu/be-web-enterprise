package categorybiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

type listCategoryBiz struct {
	store categorystore.CategoryStore
}

func NewListCategoryBiz(store categorystore.CategoryStore) *listCategoryBiz {
	return &listCategoryBiz{
		store: store,
	}
}

func (biz *listCategoryBiz) ListCategoryBiz(ctx context.Context, paging *common.Paging, filter *categorymodel.Filter) ([]categorymodel.Category, error) {
	result, err := biz.store.ListCategory(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return result, nil
}

func (biz *listCategoryBiz) ListCategoryBizForStaff(ctx context.Context) ([]categorymodel.Category, error) {
	result, err := biz.store.ListAllCategory(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return result, nil
}
