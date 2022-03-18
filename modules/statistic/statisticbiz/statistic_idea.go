package statisticbiz

import (
	"context"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticmodel"
)

type statisticIdeaBiz struct {
	ideaStore ideastore.IdeaStore
}

func NewStatisticIdeaBiz(ideaStore ideastore.IdeaStore) *statisticIdeaBiz {
	return &statisticIdeaBiz{
		ideaStore: ideaStore,
	}
}

func (biz *statisticIdeaBiz) StatisticIdeaBiz(ctx context.Context, data *statisticmodel.StatisticReq) (*statisticmodel.StatisticRespIdea, error) {
	var condition map[string]interface{}
	if data.AcaYearId != 0 {
		condition["aca_year_id"] = data.AcaYearId
	}
	if data.DepartmentId != 0 {
		condition["department_id"] = data.DepartmentId
	}

	ideas, err := biz.ideaStore.ListALlIdea(ctx, condition)
	if err != nil {
		return nil, err
	}

	var result statisticmodel.StatisticRespIdea
	result.Title = make([]string, len(ideas))
	result.LikeCount = make([]int, len(ideas))
	result.Id = make([]int, len(ideas))
	result.DislikeCount = make([]int, len(ideas))
	result.ViewCount = make([]int, len(ideas))
	result.CommentCount = make([]int, len(ideas))
	for i, _ := range ideas {
		result.Title[i] = ideas[i].Title
		result.Id[i] = ideas[i].Id
		result.LikeCount[i] = ideas[i].LikesCount
		result.DislikeCount[i] = ideas[i].DislikesCount
		result.ViewCount[i] = ideas[i].ViewsCount
		result.CommentCount[i] = ideas[i].CommentsCount
	}

	return &result, nil
}
