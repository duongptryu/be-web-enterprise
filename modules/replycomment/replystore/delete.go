package replystore

import (
	"context"
	"web/common"
	"web/modules/replycomment/replymodel"
)

func (s *sqlStore) SoftDeleteReply(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(replymodel.Reply{}.TableName()).Where("id = ?", id).Update("status", false).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DeleteReply(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(replymodel.Reply{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
