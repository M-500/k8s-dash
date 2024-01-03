package routers

import (
	"gin-server/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/2 17:12
//

func SetUpRouters(logger *zap.Logger) *gin.Engine {
	r := gin.Default()
	// 全局注册的中间件
	r.Use(middlewares.CorsMiddleWare())         // 全局处理跨域
	r.Use(middlewares.LoggerMiddleware(logger)) // 全局处理日志
	// 注册路由绑定方法
	RegisterUserRouter(r) // 注册用户模块的路由
	return r
}
