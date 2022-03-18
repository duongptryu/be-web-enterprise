package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"web/components/config"
)

func main() {
	runService()
}

func runService() {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Fatalln(err)
	}
	time.Local = loc
	//init config
	appConfig, err := config.NewAppConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	//config mode
	if appConfig.Server.ModeRun == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	f := setupLog(appConfig)
	defer f.Close()

	appCtx := setupAppContext(appConfig)

	r := gin.Default()
	r.Use(cors.Default())

	setupRouter(r, appCtx)

	r.Run(":" + appConfig.Server.Port)

}
