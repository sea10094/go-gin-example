package main

import (
	"fmt"
	"log"
	//"net/http"
	"syscall"

	"github.com/gin-gonic/gin"
        "github.com/robfig/cron"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/command_service"
	"github.com/fvbock/endless"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()

        commands, err := models.GetCommands()
        if(err != nil) {
	     fmt.Println(err.Error())
        }

        c := cron.New()
        Command := command_service.Command{}
        for _, command := range commands{
	    c.AddFunc(command.ScanPeriod, func() {
                fmt.Println(command.AgentId)
		Command.StartScan(command.AgentId, command.ScanNet, command.ConfigId)
	    })
        }
	go c.Start()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	log.Printf("[info] start http server listening %s", endPoint)

	//server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
	log.Printf("Actual pid is %d", syscall.Getpid())
	}
        err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
