package main

import (
	"flag"
	"fmt"
	"gin-server/app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

var configFile = flag.String("f", "dev.yaml", "the config file")

func main() {
	// 1. 初始化配置文件
	flag.Parse()
	serverCfg := config.MustLoadCfg(*configFile, "YML")
	fmt.Println(serverCfg.Name, serverCfg.DB.Host)
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "哈哈哈",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", serverCfg.Host, serverCfg.Port))
	//svc.SetUpServiceContext(serverCfg)
	//
	//r := routers.SetUpRouters()
	//panic(r.Run(fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.Port)))
}
