//@Author: wulinlin
//@Description: 跨域处理相关中间件
//@File:  cors
//@Version: 1.0.0
//@Date: 2023/12/01 22:57

package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		defer func() {
			if err := recover(); err != any(nil) {
				fmt.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}
