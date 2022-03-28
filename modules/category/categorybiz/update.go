package categorybiz

import (
	"context"
	log "github.com/sirupsen/logrus"
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

	if err := biz.store.UpdateCategory(ctx, cateId, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	go biz.updateTagDepartment(ctx, cateId)

	return nil
}

func (biz *updateCategoryBiz) updateTagDepartment(ctx context.Context, id int) {
	data, err := biz.store.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		log.Error(err)
		return
	}
	updateCategory := categorymodel.CategoryUpdate{
		Tags: data.SetTags(),
	}
	if err := biz.store.UpdateCategory(ctx, id, &updateCategory); err != nil {
		log.Error(err)
		return
	}
	return
}