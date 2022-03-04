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
	DecreaseLikeCountIdea(ctx, appCtx)
	IncreaseCommentCountIdea(ctx, appCtx)
	IncreaseLikeCountIdea(ctx, appCtx)
	IncreaseViewCountIdea(ctx, appCtx)
}
