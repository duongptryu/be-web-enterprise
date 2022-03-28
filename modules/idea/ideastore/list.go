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
		if v.CategoryId != 0 {
			db = db.Where("category_id = ?", v.CategoryId)
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
		if v.Search != "" {
			db = db.Where("tags LIKE ?", "%"+v.Search+"%")
		}
		if v.Order != "" {
			switch v.Order {
			case "like_desc":
				db = db.Order("likes_count desc")
			case "dislike_desc":
				db = db.Order("dislikes_count desc")
			case "view_desc":
				db = db.Order("views_count desc")
			case "created_desc":
				db = db.Order("created_at desc")
			case "like_asc":
				db = db.Order("likes_count")
			case "dislike_asc":
				db = db.Order("dislikes_count")
			case "view_asc":
				db = db.Order("views_count")
			case "created_asc":
				db = db.Order("created_at")
			}
		} else {
			db = db.Order("id desc")
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

	if err := db.Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
