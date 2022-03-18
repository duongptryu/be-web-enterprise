package commentstore

import (
	"context"
	"web/common"
	"web/modules/comment/commentmodel"
)

func (s *sqlStore) ListComment(ctx context.Context,
	condition map[string]interface{},
	filter *commentmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]commentmodel.Comment, error) {
	var result []commentmodel.Comment

	db := s.db

	db = db.Table(commentmodel.Comment{}.TableName()).Where(condition)

	// if v := filter; v != nil {

	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if paging.FakeCursor > 0 {
		db = db.Where("id < ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}

func (s *sqlStore) CountUserComment(ctx context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (map[int]int, error) {
	var result = make(map[int]int)

	type sqlData struct {
		UserId       int `gorm:"column:user_id"`
		CommentCount int `gorm:"column:count"`
	}

	var commentData []sqlData
	db := s.db

	db = db.Table(commentmodel.Comment{}.TableName()).Where(condition)

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Select("user_id, count(user_id) as count").Group("user_id").Find(&commentData).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, v := range commentData {
		result[v.UserId] = v.CommentCount
	}
	return result, nil
}
