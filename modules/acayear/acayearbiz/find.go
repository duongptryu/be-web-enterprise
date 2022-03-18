package acayearbiz

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type findAcaYear struct {
	store acayearstore.AcademicYearStore
}

func NewFindAcaYear(store acayearstore.AcademicYearStore) *findAcaYear {
	return &findAcaYear{
		store: store,
	}
}

func (biz *findAcaYear) FindAcaYear(ctx context.Context) (*acayearmodel.AcademicYear, error) {
	result, err := biz.store.FindAcaYear(ctx, map[string]interface{}{"status": true})
	if err != nil {
		return nil, common.ErrCannotListEntity(acayearmodel.EntityName, err)
	}

	return result, nil
}
