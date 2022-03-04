package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
)

func IncreaseViewCountIdea(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserViewIdea)

	ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()

		msg := <-e

		viewData := msg.Data().(CastingIdea)

		err := ideaStore.IncreaseViewCountIdea(ctx, viewData.GetIdeaId())
		if err != nil {
			log.Error(err)
		}
	}()
}
