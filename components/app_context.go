package component

import (
	"gorm.io/gorm"
	"web/components/config"
	"web/components/mailprovider"
	"web/components/mycache"
	"web/components/tokenprovider"
	"web/pubsub"
)

type AppContext interface {
	GetAppConfig() *config.AppConfig
	GetDatabase() *gorm.DB
	GetMyCache() mycache.Cache
	GetTokenProvider() tokenprovider.TokenProvider
	GetPubSub() pubsub.PubSub
	GetMailProvider() mailprovider.MailProvider
}

type appCtx struct {
	appConfig     *config.AppConfig
	database      *gorm.DB
	myCache       mycache.Cache
	tokenProvider tokenprovider.TokenProvider
	pubSub        pubsub.PubSub
	mailProvider  mailprovider.MailProvider
}

func NewAppContext(appConfig *config.AppConfig, database *gorm.DB, myCache mycache.Cache, tokenProvider tokenprovider.TokenProvider, pubSub pubsub.PubSub, mailProvider mailprovider.MailProvider) *appCtx {
	return &appCtx{
		appConfig:     appConfig,
		database:      database,
		myCache:       myCache,
		tokenProvider: tokenProvider,
		pubSub:        pubSub,
		mailProvider:  mailProvider,
	}
}

func (ctx *appCtx) GetAppConfig() *config.AppConfig {
	return ctx.appConfig
}

func (ctx *appCtx) GetDatabase() *gorm.DB {
	return ctx.database
}

func (ctx *appCtx) GetMyCache() mycache.Cache {
	return ctx.myCache
}

func (ctx *appCtx) GetTokenProvider() tokenprovider.TokenProvider {
	return ctx.tokenProvider
}

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.pubSub
}

func (ctx *appCtx) GetMailProvider() mailprovider.MailProvider {
	return ctx.mailProvider
}
