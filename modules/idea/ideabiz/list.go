package ideabiz

import (
	"context"
	"web/common"
	"web/modules/idea/ideamodel"
	"web/modules/idea/ideastore"
	"web/modules/idealikeview/idealikeviewstore"
)

type listIdeaBiz struct {
	store     ideastore.IdeaStore
	likeStore idealikeviewstore.UserLikeViewIdeaStore
}

func NewListIdeaBiz(store ideastore.IdeaStore, likeStore idealikeviewstore.UserLikeViewIdeaStore) *listIdeaBiz {
	return &listIdeaBiz{
		store:     store,
		likeStore: likeStore,
	}
}

func (biz *listIdeaBiz) ListIdeaBiz(ctx context.Context, userId int, paging *common.Paging, filter *ideamodel.Filter) ([]ideamodel.Idea, error) {
	result, err := biz.store.ListIdea(ctx, nil, filter, paging, "User", "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	//ideaIds := make([]int, len(result))
	//for i, _ := range result {
	//	ideaIds[i] = result[i].Id
	//}
	//
	//likeData, err := biz.likeStore.ListIdeaUserLike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	//if err != nil {
	//	return nil, err
	//}
	//
	//dislikeData, err := biz.likeStore.ListIdeaUserDislike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	//if err != nil {
	//	return nil, err
	//}

	for i, _ := range result {
		//if _, existLike := likeData[result[i].Id]; existLike {
		//	result[i].IsLike = true
		//} else if _, existDislike := dislikeData[result[i].Id]; existDislike {
		//	result[i].IsDislike = true
		//}
		//
		//if result[i].IsAnonymous {
		//	result[i].UserId = 0
		//	result[i].User = nil
		//}
		if i == len(result)-1 {
			paging.NextCursor = result[i].Id
		}
	}

	return result, nil
}

func (biz *listIdeaBiz) ListIdeaBizForStaff(ctx context.Context, userId int, paging *common.Paging, filter *ideamodel.Filter) ([]ideamodel.Idea, error) {
	result, err := biz.store.ListIdea(ctx, map[string]interface{}{"ideas.status": true}, filter, paging, "User", "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	//ideaIds := make([]int, len(result))
	//for i, _ := range result {
	//	ideaIds[i] = result[i].Id
	//}
	//
	//likeData, err := biz.likeStore.ListIdeaUserLike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	//if err != nil {
	//	return nil, err
	//}
	//
	//dislikeData, err := biz.likeStore.ListIdeaUserDislike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	//if err != nil {
	//	return nil, err
	//}

	for i, _ := range result {
		//if _, existLike := likeData[result[i].Id]; existLike {
		//	result[i].IsLike = true
		//} else if _, existDislike := dislikeData[result[i].Id]; existDislike {
		//	result[i].IsDislike = true
		//}
		//
		if result[i].IsAnonymous {
			result[i].UserId = 0
			result[i].User = nil
		}
		if i == len(result)-1 {
			paging.NextCursor = result[i].Id
		}
	}

	return result, nil
}

func (biz *listIdeaBiz) ListIdeaForOwner(ctx context.Context, userId int, paging *common.Paging, filter *ideamodel.Filter) ([]ideamodel.Idea, error) {
	result, err := biz.store.ListIdea(ctx, map[string]interface{}{"ideas.status": true, "ideas.user_id": userId}, filter, paging, "Category")
	if err != nil {
		return nil, common.ErrCannotListEntity(ideamodel.EntityName, err)
	}

	ideaIds := make([]int, len(result))
	for i, _ := range result {
		ideaIds[i] = result[i].Id
	}

	likeData, err := biz.likeStore.ListIdeaUserLike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	if err != nil {
		return nil, err
	}

	dislikeData, err := biz.likeStore.ListIdeaUserDislike(ctx, map[string]interface{}{"user_id": userId, "idea_id": ideaIds})
	if err != nil {
		return nil, err
	}

	for i, _ := range result {
		if _, existLike := likeData[result[i].Id]; existLike {
			result[i].IsLike = true
		} else if _, existDislike := dislikeData[result[i].Id]; existDislike {
			result[i].IsDislike = true
		}

		if result[i].IsAnonymous {
			result[i].UserId = 0
			result[i].User = nil
		}
		if i == len(result)-1 {
			paging.NextCursor = result[i].Id
		}
	}

	return result, nil
}
