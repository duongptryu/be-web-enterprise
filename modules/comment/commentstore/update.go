package commentstore

import (
	"context"
	"gorm.io/gorm"
	"web/common"
	"web/modules/comment/commentmodel"
)

func (s *sqlStore) IncreaseReplyCountComment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).Where("id = ?", id).Update("replies_count", gorm.Expr("replies_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DecreaseReplyCountComment(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(commentmodel.Comment{}.TableName()).Where("id = ?", id).Update("replies_count", gorm.Expr("replies_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}