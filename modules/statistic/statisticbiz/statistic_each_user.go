package statisticbiz

import (
	"context"
	"web/common"
	"web/modules/comment/commentstore"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewstore"
	"web/modules/statistic/statisticmodel"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
)

type statisticEachUserBiz struct {
	userStore         userstore.UserStore
	ideaLikeViewStore idealikeviewstore.UserLikeViewIdeaStore
	commentStore      commentstore.CommentStore
	ideaStore         ideastore.IdeaStore
}

func NewStatisticEachUserBiz(ideaLikeViewStore idealikeviewstore.UserLikeViewIdeaStore, userStore userstore.UserStore, commentStore commentstore.CommentStore, ideaStore ideastore.IdeaStore) *statisticEachUserBiz {
	return &statisticEachUserBiz{
		ideaLikeViewStore: ideaLikeViewStore,
		userStore:         userStore,
		commentStore:      commentStore,
		ideaStore:         ideaStore,
	}
}

func (biz *statisticEachUserBiz) StatisticEachUserBiz(ctx context.Context, userId int) (*statisticmodel.StatisticEachUser, error) {
	user, err := biz.userStore.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, common.ErrDataNotFound(usermodel.EntityName)
	}

	likeData, err := biz.ideaLikeViewStore.CountUserLike(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	dislikeData, err := biz.ideaLikeViewStore.CountUserDisLike(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	commentData, err := biz.commentStore.CountUserComment(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	viewData, err := biz.ideaLikeViewStore.CountUserViewIdea(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	ideaData, err := biz.ideaStore.CountUserPostIdea(ctx, map[string]interface{}{"user_id": userId})
	if err != nil {
		return nil, err
	}

	result := statisticmodel.StatisticEachUser{
		LikeCount:     likeData[userId],
		DislikeCount:  dislikeData[userId],
		CommentCount:  commentData[userId],
		ViewCount:     viewData[userId],
		PostIdeaCount: ideaData[userId],
	}

	return &result, nil
}
