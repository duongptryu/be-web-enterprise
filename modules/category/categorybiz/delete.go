package categorybiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

type deleteCategoryBiz struct {
	store categorystore.CategoryStore
}

func NewDeleteCategoryBiz(store categorystore.CategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{
		store: store,
	}
}

func (biz *deleteCategoryBiz) DeleteCategoryBiz(ctx context.Context, cateId int) error {
	cateDB, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": cateId})
	if err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}
	if cateDB.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	//create user in db
	if err := biz.store.DeleteCategory(ctx, cateId); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
