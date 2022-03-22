package exportbiz

import (
	"context"
	"fmt"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
)

type exportDocsZip struct {
	acaYearStore acayearstore.AcademicYearStore
}

func NewExportDocsZip(acaYearStore acayearstore.AcademicYearStore) *exportDocsZip {
	return &exportDocsZip{
		acaYearStore: acaYearStore,
	}
}

func (biz *exportDocsZip) ExportDocsZip(ctx context.Context, acaYearId int) (string, error) {
	acaYearExist, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"id": acaYearId})
	if err != nil {
		return "", err
	}
	if acaYearExist.Id == 0 {
		return "", common.ErrDataNotFound(acayearmodel.EntityName)
	}

	return fmt.Sprintf("./assets/%v/", acaYearExist.Id), nil
}
