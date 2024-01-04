package routers

import (
	"gin-server/app/controller"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/4 10:22
//

func RegisterK8sRouter(r *gin.Engine) {
	namespaceController := controller.NewNamespaceController()
	group := r.Group("api/v1/k8s")
	group.GET("/namespace", namespaceController.GetNameSpaceList)
}
