package replystore

import (
	"context"
	"web/common"
	"web/modules/replycomment/replymodel"
)

func (s *sqlStore) CreateReply(ctx context.Context, data *replymodel.ReplyCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
