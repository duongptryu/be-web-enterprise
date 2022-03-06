package ideabiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearstore"
	"web/modules/category/categorymodel"
	"web/modules/category/categorystore"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
)

type createIdeaBiz struct {
	store         ideastore.IdeaStore
	categoryStore categorystore.CategoryStore
	acaYearStore  acayearstore.AcademicYearStore
}

func NewCreateIdeaBiz(store ideastore.IdeaStore, categoryStore categorystore.CategoryStore, acaYearStore acayearstore.AcademicYearStore) *createIdeaBiz {
	return &createIdeaBiz{
		store:         store,
		categoryStore: categoryStore,
		acaYearStore:  acaYearStore,
	}
}

func (biz *createIdeaBiz) CreateIdeaBiz(ctx context.Context, data *ideamodel.IdeaCreate) error {
	cateExist, err := biz.categoryStore.FindCategory(ctx, map[string]interface{}{"id": data.CategoryId})
	if err != nil {
		return err
	}
	if cateExist.Id == 0 {
		return common.ErrDataNotFound(categorymodel.EntityName)
	}

	acaExist, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return err
	}
	if acaExist.Id == 0 || acaExist.Status == false {
		return ideamodel.ErrAcademicYearNotReady
	}

	data.AcaYearId = acaExist.Id
	data.Status = true
	if err := biz.store.CreateIdea(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
