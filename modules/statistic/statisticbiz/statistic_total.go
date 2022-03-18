package statisticbiz

import (
	"context"
	"web/modules/idea/ideastore"
	"web/modules/statistic/statisticmodel"
	"web/modules/user/userstore"
)

type statisticTotalBiz struct {
	ideaStore ideastore.IdeaStore
	userStore userstore.UserStore
}

func NewStatisticTotalBiz(ideaStore ideastore.IdeaStore, userStore userstore.UserStore) *statisticTotalBiz {
	return &statisticTotalBiz{
		ideaStore: ideaStore,
		userStore: userStore,
	}
}

func (biz *statisticTotalBiz) StatisticTotalBiz(ctx context.Context) (*statisticmodel.StatisticRespTotal, error) {
	ideas, err := biz.ideaStore.ListALlIdea(ctx, nil)
	if err != nil {
		return nil, err
	}

	var totalInteractive int
	var totalComment int

	for i, _ := range ideas {
		totalInteractive += ideas[i].LikesCount + ideas[i].DislikesCount
		totalComment += ideas[i].CommentsCount
	}

	userCount, err := biz.userStore.CountUser(ctx, nil)
	if err != nil {
		return nil, err
	}

	var result statisticmodel.StatisticRespTotal

	result.TotalComment = totalComment
	result.TotalInteractive = totalInteractive
	result.TotalUser = userCount
	result.TotalIdea = len(ideas)

	return &result, nil
}
