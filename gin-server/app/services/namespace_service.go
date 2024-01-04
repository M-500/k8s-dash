package services

import (
	"context"
	"gin-server/app/result"
	"gin-server/pkg/utils/k8s_helper"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/4 10:11
//

type NameSpaceService struct {
}

func NewNamespaceService() *NameSpaceService {
	return &NameSpaceService{}
}

func (*NameSpaceService) GetNamespaceList(ctx context.Context) ([]result.NameSpace, error) {
	k8sCfg := k8s_helper.GetK8sConn()
	list, err := k8sCfg.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	namespaceList := make([]result.NameSpace, 0)
	for _, item := range list.Items {
		namespaceList = append(namespaceList, result.NameSpace{
			Name:              item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:            string(item.Status.Phase),
		})
	}
	return namespaceList, nil
}
