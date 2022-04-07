package statisticbiz

import (
	"context"
	"time"
	"web/common"
	"web/modules/acayear/acayearmodel"
	"web/modules/acayear/acayearstore"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticmodel"
)

type statisticIdeaBiz struct {
	ideaStore    ideastore.IdeaStore
	acaYearStore acayearstore.AcademicYearStore
}

func NewStatisticIdeaBiz(ideaStore ideastore.IdeaStore, acaYearStore acayearstore.AcademicYearStore) *statisticIdeaBiz {
	return &statisticIdeaBiz{
		ideaStore:    ideaStore,
		acaYearStore: acaYearStore,
	}
}

func (biz *statisticIdeaBiz) StatisticIdeaBiz(ctx context.Context, data *statisticmodel.StatisticReq) (*statisticmodel.StatisticRespIdea, error) {
	var condition = make(map[string]interface{})
	if data.AcaYearId != 0 {
		condition["aca_year_id"] = data.AcaYearId
	}
	if data.DepartmentId != 0 {
		condition["department_id"] = data.DepartmentId
	}

	ideas, err := biz.ideaStore.ListALlIdea(ctx, condition, nil)
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

func (biz *statisticIdeaBiz) StatisticIdeaByDayBiz(ctx context.Context, acaYearId int) (*statisticmodel.StatisticRespIdeaByDay, error) {
	acaYear, err := biz.acaYearStore.FindAcaYear(ctx, map[string]interface{}{"id": acaYearId})
	if err != nil {
		return nil, err
	}
	if acaYear.Id == 0 {
		return nil, common.ErrDataNotFound(acayearmodel.EntityName)
	}

	var filter ideamodel.Filter
	filter.CreatedAtGt = &acaYear.CreatedAt

	ideas, err := biz.ideaStore.ListALlIdea(ctx, map[string]interface{}{"aca_year_id": acaYear.Id}, &filter)
	if err != nil {
		return nil, err
	}

	var ideaCount []int
	var days []string
	var check = make(map[string]int)

	for _, v := range ideas {
		var key = v.CreatedAt.Format("2006-01-02")
		if _, exist := check[key]; exist {
			check[key] = check[key] + 1
		} else {
			check[key] = 1
		}
	}

	for k, v := range check {
		ideaCount = append(ideaCount, v)
		days = append(days, k)
	}

	for d := acaYear.CreatedAt; d.After(time.Now()) == false; d = d.AddDate(0, 0, 1) {
		k := d.Format("2006-01-02")
		if v, exist := check[k]; exist {
			ideaCount = append(ideaCount, v)
			days = append(days, k)
		} else {
			ideaCount = append(ideaCount, 0)
			days = append(days, k)
		}
	}

	return &statisticmodel.StatisticRespIdeaByDay{
		ideaCount,
		days,
	}, nil
}
