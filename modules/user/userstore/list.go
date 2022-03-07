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
