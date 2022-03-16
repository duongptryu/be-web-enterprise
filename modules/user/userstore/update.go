package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
)

func (s *sqlStore) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}


func (s *sqlStore) UpdateUserSelf(ctx context.Context, id int, data *usermodel.UserUpdateSelf) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
