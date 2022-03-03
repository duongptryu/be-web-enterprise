package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*usermodel.User, error) {
	var result usermodel.User

	db := s.db

	db = db.Table(usermodel.User{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
