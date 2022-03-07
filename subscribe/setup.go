package subscribe

import (
	"context"
	component "web/components"
)

type CastingIdea interface {
	GetIdeaId() int
}

func SetupSubscribe(appCtx component.AppContext) {
	ctx := context.Background()
	IncreaseLikeCountIdea(ctx, appCtx)
	DecreaseLikeCountIdea(ctx, appCtx)
	IncreaseDisLikeCountIdea(ctx, appCtx)
	DecreaseDisLikeCountIdea(ctx, appCtx)
	IncreaseCommentCountIdea(ctx, appCtx)
	DecreaseCommentCountIdea(ctx, appCtx)
	IncreaseViewCountIdea(ctx, appCtx)

	DecreaseReplyCountComment(ctx, appCtx)
	IncreaseReplyCountComment(ctx, appCtx)
}
