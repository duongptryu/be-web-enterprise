package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
	"web/common"
	"web/components"
	"web/components/config"
	"web/components/mailprovider/sendgridprovider"
	"web/components/mycache"
	"web/components/tokenprovider/jwt"
	"web/modules/user/usermodel"
	"web/modules/user/userstore"
	"web/pubsub/pubsublocal"
	"web/subscribe"
)

func createDsnDb(username, password, host, port, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
}

func setupAppContext(appConfig *config.AppConfig) component.AppContext {
	//init database
	databaseDsn := createDsnDb(appConfig.Database.Username, appConfig.Database.Password, appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.DatabaseName)
	FDDatabase, err := gorm.Open(mysql.Open(databaseDsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database notification- ", err)
	}
	FDDatabase = FDDatabase.Debug()

	//init cache
	myCache := mycache.NewMyCache()

	//init token provider
	tokenProvider := jwt.NewJwtProvider(appConfig.Token)

	//init upload provider
	//s3Provider := uploadprovider.NewS3Provider(appConfig.S3AWS.BucketName, appConfig.S3AWS.Region, appConfig.S3AWS.ApiKey, appConfig.S3AWS.Secret, appConfig.S3AWS.Domain)

	//init pubsub local
	pubSubLocal := pubsublocal.NewPubSub()

	//init sendgrid provider
	sendgridMail := sendgridprovider.NewSendGridProvider(appConfig.SendgridSecretKey)
	//init smtpProvider
	//smtpProvider := smtp.NewSmtpProvider(appConfig.Smtp.Email, appConfig.Smtp.Password)

	//init app context
	appCtx := component.NewAppContext(appConfig, FDDatabase, myCache, tokenProvider, pubSubLocal, sendgridMail)

	//setup subscribe
	subscribe.SetupSubscribe(appCtx)

	//init admin account
	InitAdminAccount(appCtx)

	return appCtx
}

func setupLog(appConfig *config.AppConfig) *os.File {
	f, err := os.OpenFile("web-log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("error opening file: %v", err)
	}
	log.SetOutput(f)
	//config log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if appConfig.Server.LevelLog >= 0 && appConfig.Server.LevelLog <= 6 {
		log.SetLevel(log.AllLevels[appConfig.Server.LevelLog])
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	return f
}

func InitAdminAccount(appCtx component.AppContext) {
	userStore := userstore.NewSQLStore(appCtx.GetDatabase())
	userDB, err := userStore.FindUser(context.Background(), map[string]interface{}{"role": common.RoleAdmin})
	if err != nil {
		log.Fatalln(err)
	}
	if userDB.Id == 0 {
		adminAccount := usermodel.UserCreate{
			Email:       "admin@gmail.com",
			Password:    "123123",
			FullName:    "administrator",
			Gender:      "Male",
			Role:        common.RoleAdmin,
			DateOfBirth: time.Now(),
			Status:      true,
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminAccount.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		adminAccount.Password = string(hashedPassword)

		if err := userStore.CreateUser(context.Background(), &adminAccount); err != nil {
			log.Fatalln(err)
		}

		log.Println("Init account admin: admin@gmail.com - 123456789")
	}
}
