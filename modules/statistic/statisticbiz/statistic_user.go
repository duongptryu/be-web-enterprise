package statisticbiz

import (
	"context"
	"web/modules/comment/commentstore"
	"web/modules/idealikeview/idealikeviewstore"
	"web/modules/statistic/statisticmodel"
	"web/modules/user/userstore"
)

type statisticUserBiz struct {
	userStore         userstore.UserStore
	ideaLikeViewStore idealikeviewstore.UserLikeViewIdeaStore
	commentStore      commentstore.CommentStore
}

func NewStatisticUserBiz(ideaLikeViewStore idealikeviewstore.UserLikeViewIdeaStore, userStore userstore.UserStore, commentStore commentstore.CommentStore) *statisticUserBiz {
	return &statisticUserBiz{
		ideaLikeViewStore: ideaLikeViewStore,
		userStore:         userStore,
		commentStore:      commentStore,
	}
}

func (biz *statisticUserBiz) StatisticUserBiz(ctx context.Context, data *statisticmodel.StatisticUser) (*statisticmodel.StatisticRespUser, error) {
	var condition = make(map[string]interface{})
	if data.DepartmentId > 0 {
		condition["department_id"] = data.DepartmentId
	} else {
		condition = nil
	}
	users, err := biz.userStore.ListUserWithoutPaging(ctx, condition)
	if err != nil {
		return nil, err
	}

	var result statisticmodel.StatisticRespUser

	result.UsersName = make([]string, len(users))
	result.UsersId = make([]int, len(users))
	result.UsersInteractive = make([]int, len(users))

	if len(users) == 0 {
		return &result, nil
	}

	listUserIds := make([]int, len(users))
	for i, _ := range users {
		listUserIds[i] = users[i].Id
	}

	likeData, err := biz.ideaLikeViewStore.CountUserLike(ctx, map[string]interface{}{"user_id": listUserIds})
	if err != nil {
		return nil, err
	}

	dislikeData, err := biz.ideaLikeViewStore.CountUserDisLike(ctx, map[string]interface{}{"user_id": listUserIds})
	if err != nil {
		return nil, err
	}

	commentData, err := biz.commentStore.CountUserComment(ctx, map[string]interface{}{"user_id": listUserIds})
	if err != nil {
		return nil, err
	}

	for i, _ := range users {
		result.UsersName[i] = users[i].FullName
		result.UsersId[i] = users[i].Id
		result.UsersInteractive[i] = likeData[users[i].Id] + dislikeData[users[i].Id] + commentData[users[i].Id]
	}

	return &result, nil
}
