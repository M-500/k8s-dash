package response

import (
	"gin-server/pkg/comm/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResult struct {
	ErrorCode int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Success(c *gin.Context) {
	res := &JsonResult{
		ErrorCode: e.SUCCESS,
		Message:   e.GetMessage(e.SUCCESS),
		Data:      nil,
		Success:   true,
	}
	c.JSON(http.StatusOK, res)
}

func JsonSuccessMsg(c *gin.Context, msg string) {
	res := &JsonResult{
		ErrorCode: e.SUCCESS,
		Message:   msg,
		Data:      nil,
		Success:   true,
	}
	c.JSON(http.StatusOK, res)
}

func JsonSuccessCode(c *gin.Context, code int) {
	res := &JsonResult{
		ErrorCode: code,
		Message:   e.GetMessage(code),
		Data:      nil,
		Success:   true,
	}
	c.JSON(http.StatusOK, res)
}

func JsonSuccessData(c *gin.Context, msg string, data any) {
	res := &JsonResult{
		ErrorCode: e.SUCCESS,
		Message:   msg,
		Data:      data,
		Success:   true,
	}
	c.JSON(http.StatusOK, res)
}

func JsonFail(c *gin.Context) {
	res := &JsonResult{
		ErrorCode: e.ERROR,
		Message:   e.GetMessage(e.ERROR),
		Data:      nil,
		Success:   false,
	}
	c.JSON(http.StatusOK, res)
}

func JsonFailMsg(c *gin.Context, msg string) {
	res := &JsonResult{
		ErrorCode: e.ERROR,
		Message:   msg,
		Data:      nil,
		Success:   false,
	}
	c.JSON(http.StatusOK, res)
}

func JsonFailCode(c *gin.Context, code int) {
	res := &JsonResult{
		ErrorCode: code,
		Message:   e.GetMessage(code),
		Data:      nil,
		Success:   false,
	}
	c.JSON(http.StatusOK, res)
}

func JsonFailData(c *gin.Context, msg string, data any) {
	res := &JsonResult{
		ErrorCode: e.ERROR,
		Message:   msg,
		Data:      data,
		Success:   false,
	}
	c.JSON(http.StatusOK, res)
}
