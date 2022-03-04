package idealikeviewstore

import (
	"context"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
)

func (s *sqlStore) DeleteUserLikeIdea(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(idealikeviewmodel.UserLikeIdea{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStore) DeleteUserDislikeIdea(ctx context.Context, condition map[string]interface{}) error {
	db := s.db

	if err := db.Table(idealikeviewmodel.UserDislikeIdea{}.TableName()).Where(condition).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
