package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/idea/ideastore"
)

func DecreaseCommentCountIdea(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicDecreaseCommentCountIdea)

	ideaStore := ideastore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-e

			ideaId := msg.Data().(int)

			err := ideaStore.DecreaseCommentCountIdea(ctx, ideaId)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
