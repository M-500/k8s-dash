package main

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/3 19:26
//

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var KubeConfigSet *kubernetes.Clientset

func main() {
	kubeConfig := "test/t1/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	KubeConfigSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ctx := context.TODO()
	lists, err := KubeConfigSet.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		return
	}
	for i, i2 := range lists.Items {
		fmt.Println(i, i2.Name)
	}
}
