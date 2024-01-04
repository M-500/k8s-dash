package controller

import (
	"context"
	"gin-server/app/services"
	"gin-server/pkg/comm/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/4 10:00
//

type NamespaceController struct {
}

func NewNamespaceController() *NamespaceController {
	return &NamespaceController{}
}

func (*NamespaceController) GetNameSpaceList(c *gin.Context) {
	ctx := context.TODO()
	service := services.NewNamespaceService()
	list, err := service.GetNamespaceList(ctx)
	if err != nil {
		response.JsonFailMsg(c, err.Error())
		return
	}
	response.JsonSuccessData(c, "成功", list)
	return
}
