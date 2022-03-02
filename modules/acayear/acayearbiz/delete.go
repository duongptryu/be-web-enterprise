package acayearbiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type deleteAcaYearBiz struct {
	store acayearstore.AcademicYearStore
}

func NewDeleteAcaYearBiz(store acayearstore.AcademicYearStore) *deleteAcaYearBiz {
	return &deleteAcaYearBiz{
		store: store,
	}
}

func (biz *deleteAcaYearBiz) DeleteAcaYearBiz(ctx context.Context, acaYearId int) error {
	userDB, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"id": acaYearId})
	if err != nil {
		return common.ErrCannotDeleteEntity(acayearmodel.EntityName, err)
	}
	if userDB.Id == 0 {
		return common.ErrDataNotFound(acayearmodel.EntityName)
	}

	//create user in db
	if err := biz.store.DeleteAcaYear(ctx, acaYearId); err != nil {
		return common.ErrCannotDeleteEntity(acayearmodel.EntityName, err)
	}

	return nil
}
