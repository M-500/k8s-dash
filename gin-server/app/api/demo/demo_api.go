//@Author: wulinlin
//@Description:
//@File:  demo_api
//@Version: 1.0.0
//@Date: 2024/01/02 21:28

package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoApi struct {
}

func (*DemoApi) DemoController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "哈哈哈",
	})
	return
}
