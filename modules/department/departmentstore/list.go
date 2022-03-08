package departmentstore

import (
	"context"
	"web/common"
	"web/modules/department/departmentmodel"
)

func (s *sqlStore) ListDepartment(ctx context.Context,
	condition map[string]interface{},
	filter *departmentmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]departmentmodel.Department, error) {
	var result []departmentmodel.Department

	db := s.db

	db = db.Table(departmentmodel.Department{}.TableName()).Where(condition)

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
