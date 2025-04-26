package main

import (
	"flag"
	"log"
	"myip/config"
	"myip/handler"

	"github.com/gin-gonic/gin"
	"github.com/xxxsen/common/logger"
	"github.com/xxxsen/common/webapi"
	"go.uber.org/zap"
)

var (
	conf = flag.String("config", "./config.json", "config file path")
)

func main() {
	flag.Parse()
	cc, err := config.Parse(*conf)
	if err != nil {
		log.Fatalf("read config file fail, path:%s, err:%v", *conf, err)
	}
	logkit := logger.Init(cc.LogConfig.File, cc.LogConfig.Level, int(cc.LogConfig.FileCount), int(cc.LogConfig.FileSize), int(cc.LogConfig.KeepDays), cc.LogConfig.Console)
	logkit.Info("read config succ", zap.Any("config", *cc))
	iph := handler.NewIPHandler(cc.Headers)
	engine, err := webapi.NewEngine("/", cc.Bind, webapi.WithRegister(func(c *gin.RouterGroup) {
		c.GET("/", iph.HandleGetIP)
	}))
	if err != nil {
		logkit.Fatal("create webapi engine fail", zap.Error(err))
	}
	if err := engine.Run(); err != nil {
		logkit.Fatal("run webapi engine fail", zap.Error(err))
	}
}
