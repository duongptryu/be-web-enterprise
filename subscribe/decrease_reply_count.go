package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/comment/commentstore"
)

func DecreaseReplyCountComment(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicDecreaseReplyCountComment)

	commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-e

			ideaId := msg.Data().(int)

			err := commentStore.DecreaseReplyCountComment(ctx, ideaId)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
