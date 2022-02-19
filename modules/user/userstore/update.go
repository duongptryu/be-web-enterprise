package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
)

func (s *sqlStore) UpdateStatusUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
