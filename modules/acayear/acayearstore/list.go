package acayearstore

import (
	"context"
	"web/common"
	"web/modules/acayear/acayearmodel"
)

func (s *sqlStore) ListAcaYear(ctx context.Context,
	condition map[string]interface{},
	filter *acayearmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]acayearmodel.AcademicYear, error) {
	var result []acayearmodel.AcademicYear

	db := s.db

	db = db.Table(acayearmodel.AcademicYear{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (s *sqlStore) ListAcaYearWithoutPaging(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) ([]acayearmodel.AcademicYear, error) {
	var result []acayearmodel.AcademicYear

	db := s.db

	db = db.Table(acayearmodel.AcademicYear{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
