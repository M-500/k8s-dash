package main

import (
	"flag"
	"fmt"
	"gin-server/app/config"
	"gin-server/pkg/utils/k8s_helper"
	"gin-server/pkg/utils/logger_helper"
	"gin-server/routers"
	"go.uber.org/zap"
)

var configFile = flag.String("f", "app/etc/dev.yaml", "the config file")
var logger *zap.Logger

func main() {
	// 1. 初始化配置文件
	flag.Parse()
	serverCfg := config.MustLoadCfg(*configFile, "YML")
	fmt.Println(serverCfg.Name, serverCfg.DB.Host)

	logger, err := logger_helper.SetUpLogger(true, &serverCfg.Logs)
	if err != nil {
		panic(err)
	}

	// 初始化K8S
	k8s_helper.SetUpK8s()
	r := routers.SetUpRouters(logger)
	panic(r.Run(fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.Port)))
}
