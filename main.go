package main

import (
	"context"
	"fmt"
	"gg_web_tmpl/cache"
	"gg_web_tmpl/common/config"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/model"
	"gg_web_tmpl/router"
)

func main() {
	// 初始化配置
	err := config.InitConfig("./conf/config.yaml")
	if err != nil {
		panic(err)
		return
	}
	// 初始化logger
	if err := log.InitLogger(config.GetConf().App.LogPath, config.GetConf().App.LogLevel); err != nil {
		panic(err)
		return
	}
	// 初始化mysql
	if err := model.InitMySQL(config.GetConf().MySQL); err != nil {
		log.GetLogger().Errorf("init mysql error: %v", err)
		return
	}
	// 初始化redis
	if err := cache.InitRedis(context.Background()); err != nil {
		log.GetLogger().Errorf("init redis error: %v", err)
		return
	}
	// 初始化gin
	r := router.SetupRouter()
	// 启动
	if err := r.Run(fmt.Sprintf(":%d", config.GetConf().App.Port)); err != nil {
		log.GetLogger().Errorf("setup router error: %v", err)
		return
	}
}
