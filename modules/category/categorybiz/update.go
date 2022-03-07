package categorybiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
)

type updateCategoryBiz struct {
	store categorystore.CategoryStore
}

func NewUpdateCategoryBiz(store categorystore.CategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{
		store: store,
	}
}

func (biz *updateCategoryBiz) UpdateCategoryBiz(ctx context.Context, cateId int, data *categorymodel.CategoryUpdate) error {
	cateDB, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": cateId})
	if err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}
	if cateDB.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	//create user in db
	if err := biz.store.UpdateCategory(ctx, cateId, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
