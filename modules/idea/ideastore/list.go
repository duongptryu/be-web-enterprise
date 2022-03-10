package ideastore

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
)

func (s *sqlStore) ListIdea(ctx context.Context,
	condition map[string]interface{},
	filter *ideamodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]ideamodel.Idea, error) {
	var result []ideamodel.Idea

	db := s.db

	db = db.Table(ideamodel.Idea{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.UserId != 0 {
			db = db.Where("user_id = ?", v.UserId)
		}
		if v.IsAnonymous != "" && (v.IsAnonymous == "false" || v.IsAnonymous == "true") {
			db = db.Where("is_anonymous = ?", v.IsAnonymous)
		}
		if v.Title != "" {
			db = db.Where("title LIKE ?", "%"+v.Title+"%")
		}
		if v.AcaYearId != 0 {
			db = db.Where("aca_year_id = ?", v.AcaYearId)
		}
		if v.DislikeGt > 0 {
			db = db.Where("dislike_count >= ?", v.DislikeGt)
		}
		if v.DislikeSt > 0 {
			db = db.Where("dislikes_count <= ?", v.DislikeSt)
		}
		if v.LikeGt > 0 {
			db = db.Where("likes_count >= ?", v.LikeGt)
		}
		if v.LikeSt > 0 {
			db = db.Where("likes_count <= ?", v.LikeSt)
		}
		if v.ViewGt > 0 {
			db = db.Where("views_count >= ?", v.ViewGt)
		}
		if v.ViewSt > 0 {
			db = db.Where("views_count <= ?", v.ViewSt)
		}
		if v.CategoryName != "" {
			db = db.Joins("JOIN categories on categories.id = ideas.category_id").Where("categories.name LIKE ?", "%"+v.CategoryName+"%")
		}
		if v.DepartmentName != "" {
			db = db.Joins("JOIN departments on departments.id = ideas.department_id").Where("departments.name LIKE ?", "%"+v.DepartmentName+"%")
		}
		if v.AcaYear != "" {
			db = db.Joins("JOIN academic_years on academic_years.id = ideas.aca_year_id").Where("academic_years.name LIKE ?", "%"+v.AcaYear+"%")
		}
	}

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
