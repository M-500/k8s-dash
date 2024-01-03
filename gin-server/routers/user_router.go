package routers

import "github.com/gin-gonic/gin"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/3 11:27
//

func RegisterUserRouter(r *gin.Engine) {
	group := r.Group("/na")
	group.POST("")
}
