package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
)

func IncreaseViewCountIdea(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicIncreaseViewCountIdea)

	ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-e

			ideaId := msg.Data().(int)

			err := ideaStore.IncreaseViewCountIdea(ctx, ideaId)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
