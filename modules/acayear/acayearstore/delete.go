package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
)

func (s *sqlStore) DeleteAcaYear(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(acayearmodel.AcademicYear{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
