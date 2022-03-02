package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
)

func (s *sqlStore) FindAcaYear(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*acayearmodel.AcademicYear, error) {
	var result acayearmodel.AcademicYear

	db := s.db

	db = db.Table(acayearmodel.AcademicYear{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
