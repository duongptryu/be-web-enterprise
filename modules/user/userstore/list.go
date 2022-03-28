package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
)

func (s *sqlStore) ListUser(ctx context.Context,
	condition map[string]interface{},
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db

	db = db.Table(usermodel.User{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.Role != "" {
			db = db.Where("role = ?", v.Role)
		}
		if v.Status != "" && (v.Status == "false" || v.Status == "true") {
			db = db.Where("status = ?", v.Status)
		}
		if v.DepartmentId != 0 {
			db = db.Where("department_id = ?", v.DepartmentId)
		}
		if v.Email != "" {
			db = db.Where("email LIKE ?", "%"+v.Email+"%")
		}
		if v.FullName != "" {
			db = db.Where("full_name LIKE ?", "%"+v.FullName+"%")
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

func (s *sqlStore) ListUserWithoutPaging(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db

	db = db.Table(usermodel.User{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (s *sqlStore) CountUser(ctx context.Context,
	condition map[string]interface{},
) (int, error) {
	db := s.db

	db = db.Table(usermodel.User{}.TableName()).Where(condition)

	var result int64

	if err := db.Count(&result).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	resultInt := int(result)

	return resultInt, nil
}
