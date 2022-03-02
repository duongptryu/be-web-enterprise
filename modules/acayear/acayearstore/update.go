package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
)

func (s *sqlStore) UpdateAcaYear(ctx context.Context, id int, data *acayearmodel.AcademicYearUpdate) error {
	db := s.db

	if err := db.Table(acayearmodel.AcademicYearUpdate{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
