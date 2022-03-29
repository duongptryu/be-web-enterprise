package statisticbiz

import (
	"context"
	"web/modules/category/categorystore"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticmodel"
)

type statisticCountIdeaInCategoryBiz struct {
	ideaStore     ideastore.IdeaStore
	categoryStore categorystore.CategoryStore
}

func NewStatisticCountIdeaInCategoryBiz(ideaStore ideastore.IdeaStore, categoryStore categorystore.CategoryStore) *statisticCountIdeaInCategoryBiz {
	return &statisticCountIdeaInCategoryBiz{
		ideaStore:     ideaStore,
		categoryStore: categoryStore,
	}
}

func (biz *statisticCountIdeaInCategoryBiz) StatisticCountIdeaInCategoryBiz(ctx context.Context) ([]statisticmodel.StatisticCountIdeaCategory, error) {
	allCategory, err := biz.categoryStore.ListAllCategory(ctx, nil)
	if err != nil {
		return nil, err
	}

	var categoryIds = make([]int, len(allCategory))
	for i, _ := range allCategory {
		categoryIds[i] = allCategory[i].Id
	}

	query, err := biz.ideaStore.CountIdeaByCategoryId(ctx, categoryIds)
	if err != nil {
		return nil, err
	}

	var result = make([]statisticmodel.StatisticCountIdeaCategory, len(allCategory))

	for i, _ := range allCategory {
		result[i].CategoryName = allCategory[i].Name
		result[i].CategoryId = allCategory[i].Id
		if count, exist := query[allCategory[i].Id]; exist {
			result[i].NumberIdea = count
		} else {
			result[i].NumberIdea = 0
		}
	}

	return result, nil
}
