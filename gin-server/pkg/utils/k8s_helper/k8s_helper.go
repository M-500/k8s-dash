package k8s_helper

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/3 19:54
import (
	myCfg "gin-server/app/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var KubeConfigSet *kubernetes.Clientset

func SetUpK8s() {
	config, err := clientcmd.BuildConfigFromFlags("", myCfg.Instance.KubeConfigPath)
	if err != nil {
		panic(err.Error())
	}
	// create the client set
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	KubeConfigSet = clientSet
}

func GetK8sConn() *kubernetes.Clientset {
	if KubeConfigSet == nil {
		SetUpK8s()
	}
	return KubeConfigSet
}
