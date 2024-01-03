package routers

import (
	"gin-server/app/controller"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/3 11:27
//

func RegisterUserRouter(r *gin.Engine) {
	userController := controller.NewUserController()
	group := r.Group("/api/v1/na")
	group.GET("/demo", userController.Demo)
}
