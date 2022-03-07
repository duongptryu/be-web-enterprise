package commentstore

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
)


func (s *sqlStore) SoftDeleteComment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).Where("id = ?", id).Update("status", false).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DeleteComment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}