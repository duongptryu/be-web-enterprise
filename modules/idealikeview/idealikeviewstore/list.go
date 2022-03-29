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

func (s *sqlStore) ListIdeaUserLike(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		IdeaId int `gorm:"column:idea_id"`
		UserId int `gorm:"column:user_id"`
	}

	var likeData []sqlData
	db := s.db

	db = db.Table(idealikeviewmodel.UserLikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select([]string{"idea_id", "user_id"}).Find(&likeData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range likeData {
		result[v.IdeaId] = v.UserId
	}
	return result, nil
}

func (s *sqlStore) ListIdeaUserDislike(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		IdeaId int `gorm:"column:idea_id"`
		UserId int `gorm:"column:user_id"`
	}

	var likeData []sqlData
	db := s.db

	db = db.Table(idealikeviewmodel.UserDislikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select([]string{"idea_id", "user_id"}).Find(&likeData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range likeData {
		result[v.IdeaId] = v.UserId
	}
	return result, nil
}

func (s *sqlStore) CountUserLike(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		UserId    int `gorm:"column:user_id"`
		LikeCount int `gorm:"column:count"`
	}

	var likeData []sqlData
	db := s.db

	db = db.Table(idealikeviewmodel.UserLikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select("user_id, count(user_id) as count").Group("user_id").Find(&likeData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range likeData {
		result[v.UserId] = v.LikeCount
	}
	return result, nil
}

func (s *sqlStore) CountUserDisLike(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		UserId       int `gorm:"column:user_id"`
		dislikeCount int `gorm:"column:count"`
	}

	var disLikeData []sqlData
	db := s.db

	db = db.Table(idealikeviewmodel.UserDislikeIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select("user_id, count(user_id) as count").Group("user_id").Find(&disLikeData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range disLikeData {
		result[v.UserId] = v.dislikeCount
	}
	return result, nil
}

func (s *sqlStore) CountUserViewIdea(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		UserId    int `gorm:"column:user_id"`
		ViewCount int `gorm:"column:count"`
	}

	var viewIdeaCount []sqlData
	db := s.db

	db = db.Table(idealikeviewmodel.UserViewIdea{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select("user_id, count(user_id) as count").Group("user_id").Find(&viewIdeaCount).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range viewIdeaCount {
		result[v.UserId] = v.ViewCount
	}
	return result, nil
}
