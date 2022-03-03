package userstore

import (
	"context"
	"web/common"
	"web/modules/user/usermodel"
)

func (s *sqlStore) SoftDeleteUser(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id = ?", id).Update("status", false).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
