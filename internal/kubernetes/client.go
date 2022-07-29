package kubernetes

import (
	"fmt"

	"github.com/rymdo/kube-watchtower/v2/internal/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

type Kubernetes struct {
	client *kubernetes.Clientset
}

func New(cfg *config.Config) Kubernetes {
	var apiConfig *rest.Config

	if cfg.KubeconfigExists {
		fmt.Printf("kubernetes: using kubeconfig mode - '%s'\n", cfg.KubeconfigPath)
		apiConfig = configKubeconfig(cfg.KubeconfigPath)
	} else {
		fmt.Println("kubernetes: using in-cluster mode")
		apiConfig = configInCluster()
	}

	clientset, err := kubernetes.NewForConfig(apiConfig)
	if err != nil {
		panic(err.Error())
	}

	_, err = apiextensionsclient.NewForConfig(apiConfig)
	if err != nil {
		panic(err)
	}

	return Kubernetes{
		client: clientset,
	}
}

func configInCluster() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func configKubeconfig(kubeconfigPath string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err.Error())
	}
	return config
}
