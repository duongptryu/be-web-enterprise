package departmentstore

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
)

func (s *sqlStore) FindDepartment(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*departmentmodel.Department, error) {
	var result departmentmodel.Department

	db := s.db

	db = db.Table(departmentmodel.Department{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
