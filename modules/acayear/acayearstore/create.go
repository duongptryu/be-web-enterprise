package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
)

func (s *sqlStore) CreateAcaYear(ctx context.Context, data *acayearmodel.AcademicYearCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
