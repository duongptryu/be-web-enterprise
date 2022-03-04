package idealikeviewstore

import (
	"context"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
)

func (s *sqlStore) FindUserViewIdea(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*idealikeviewmodel.UserViewIdea, error) {
	var result idealikeviewmodel.UserViewIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserViewIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}

func (s *sqlStore) FindUserLikeIdea(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*idealikeviewmodel.UserLikeIdea, error) {
	var result idealikeviewmodel.UserLikeIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserLikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}

func (s *sqlStore) FindUserDislikeIdea(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*idealikeviewmodel.UserDislikeIdea, error) {
	var result idealikeviewmodel.UserDislikeIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserDislikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
