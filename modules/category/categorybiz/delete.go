package categorybiz

import (
	"context"
	"web/common"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
	"web/modules/idea/ideastore"
)

type deleteCategoryBiz struct {
	store     categorystore.CategoryStore
	ideaStore ideastore.IdeaStore
}

func NewDeleteCategoryBiz(store categorystore.CategoryStore, ideaStore ideastore.IdeaStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{
		store:     store,
		ideaStore: ideaStore,
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

	idea, err := biz.ideaStore.FindIdea(ctx, map[string]interface{}{"category_id": cateDB.Id})
	if err != nil {
		return err
	}

	if idea.Id != 0 {
		return categorymodel.ErrCannotDelCategory
	}

	//create user in db
	if err := biz.store.DeleteCategory(ctx, cateId); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
