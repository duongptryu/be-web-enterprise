package replystore

import (
	"context"
	"web/common"
	"web/modules/replycomment/replymodel"
)

func (s *sqlStore) FindReply(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*replymodel.Reply, error) {
	var result replymodel.Reply

	db := s.db

	db = db.Table(replymodel.Reply{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
