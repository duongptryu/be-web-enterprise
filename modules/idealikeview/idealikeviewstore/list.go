package idealikeviewstore

import (
	"context"
	"web/common"
	"web/modules/idealikeview/idealikeviewmodel"
)

func (s *sqlStore) ListUserLikeIdea(ctx context.Context,
	condition map[string]interface{},
	filter *idealikeviewmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]idealikeviewmodel.UserLikeIdea, error) {
	var result []idealikeviewmodel.UserLikeIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserLikeIdea{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (s *sqlStore) ListUserDislikeIdea(ctx context.Context,
	condition map[string]interface{},
	filter *idealikeviewmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]idealikeviewmodel.UserDislikeIdea, error) {
	var result []idealikeviewmodel.UserDislikeIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserDislikeIdea{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (s *sqlStore) ListUserViewIdea(ctx context.Context,
	condition map[string]interface{},
	filter *idealikeviewmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]idealikeviewmodel.UserViewIdea, error) {
	var result []idealikeviewmodel.UserViewIdea

	db := s.db

	db = db.Table(idealikeviewmodel.UserViewIdea{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("created_at desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
