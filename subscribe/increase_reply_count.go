package subscribe

import (
	"context"
	log "github.com/sirupsen/logrus"
	"web/common"
	component "web/components"
	"web/modules/comment/commentstore"
)

func IncreaseReplyCountComment(ctx context.Context, appCtx component.AppContext) {
	e, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicIncreaseReplyCountComment)

	commentStore := commentstore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-e

			ideaId := msg.Data().(int)

			err := commentStore.IncreaseReplyCountComment(ctx, ideaId)
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
