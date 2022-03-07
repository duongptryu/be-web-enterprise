package acayearbiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type createAcaYearBiz struct {
	store acayearstore.AcademicYearStore
}

func NewCreateAcaYearBiz(store acayearstore.AcademicYearStore) *createAcaYearBiz {
	return &createAcaYearBiz{
		store: store,
	}
}

func (biz *createAcaYearBiz) CreateAcaYearBiz(ctx context.Context, data *acayearmodel.AcademicYearCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return err
	}
	if result.Id != 0 {
		return acayearmodel.ErrOverlapAcaYear
	}

	if err := biz.store.CreateAcaYear(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(acayearmodel.EntityName, err)
	}

	return nil
}
