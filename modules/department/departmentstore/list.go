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

	if v := filter; v != nil {
		if v.LeaderName != "" {
			db = db.Joins("JOIN users on departments.id = user.department_id AND users.role = 'QAC'").Where("users.name LIKE ?", "%"+v.LeaderName+"%")
		}
		if v.Name != "" {
			db = db.Where("name LIKE ?", "%"+v.Name+"%")
		}
		if v.Search != "" {
			db = db.Where("tags LIKE ?", "%"+v.Search+"%")
		}
	}

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

func (s *sqlStore) ListDepartmentForStaff(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) ([]departmentmodel.Department, error) {
	var result []departmentmodel.Department

	db := s.db

	db = db.Table(departmentmodel.Department{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
