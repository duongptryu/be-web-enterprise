package acayearbiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type updateAcaYearBiz struct {
	store acayearstore.AcademicYearStore
}

func NewUpdateUserBiz(store acayearstore.AcademicYearStore) *updateAcaYearBiz {
	return &updateAcaYearBiz{
		store: store,
	}
}

func (biz *updateAcaYearBiz) UpdateUserBiz(ctx context.Context, acaYearId int, data *acayearmodel.AcademicYearUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	acaYearDB, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"id": acaYearId})
	if err != nil {
		return common.ErrCannotUpdateEntity(acayearmodel.EntityName, err)
	}
	if acaYearDB.Id == 0 {
		return common.ErrDataNotFound(acayearmodel.EntityName)
	}

	//create user in db
	if err := biz.store.UpdateAcaYear(ctx, acaYearId, data); err != nil {
		return common.ErrCannotUpdateEntity(acayearmodel.EntityName, err)
	}

	return nil
}
