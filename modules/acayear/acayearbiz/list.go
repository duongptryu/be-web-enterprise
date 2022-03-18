package acayearbiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type listAcaYearBiz struct {
	store acayearstore.AcademicYearStore
}

func NewListAcaYearBiz(store acayearstore.AcademicYearStore) *listAcaYearBiz {
	return &listAcaYearBiz{
		store: store,
	}
}

func (biz *listAcaYearBiz) ListAcaYearBiz(ctx context.Context, paging *common.Paging, filter *acayearmodel.Filter) ([]acayearmodel.AcademicYear, error) {
	result, err := biz.store.ListAcaYear(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(acayearmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listAcaYearBiz) ListAcaYearBizWithoutPaging(ctx context.Context) ([]acayearmodel.AcademicYear, error) {
	result, err := biz.store.ListAcaYearWithoutPaging(ctx, nil)
	if err != nil {
		return nil, common.ErrCannotListEntity(acayearmodel.EntityName, err)
	}

	return result, nil
}
