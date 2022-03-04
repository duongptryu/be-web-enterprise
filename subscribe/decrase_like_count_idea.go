package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
)

func DecreaseLikeCountIdea(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicStaffDislikeIdea)

	ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()

		msg := <-e

		ideaId := msg.Data().(int)

		err := ideaStore.DecreaseLikeCountIdea(ctx, ideaId)
		if err != nil {
			log.Error(err)
		}
	}()
}
