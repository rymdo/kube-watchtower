package kubernetes

import (
	"time"

	"github.com/rymdo/kube-watchtower/v2/internal/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

type Kubernetes struct {
	logger              *utils.Logger
	cfg                 *utils.Config
	client              *kubernetes.Clientset
	clientApiExtensions *apiextensionsclient.Clientset
}

func New(logger *utils.Logger, cfg *utils.Config) Kubernetes {
	var apiConfig *rest.Config

	if cfg.KubeconfigExists {
		logger.Infof("kubernetes: using kubeconfig mode - '%s'", cfg.KubeconfigPath)
		apiConfig = configKubeconfig(cfg.KubeconfigPath)
	} else {
		logger.Infof("kubernetes: using in-cluster mode")
		apiConfig = configInCluster()
	}

	apiConfig.Timeout = time.Second * 5

	clientset, err := kubernetes.NewForConfig(apiConfig)
	if err != nil {
		panic(err.Error())
	}

	clientApiExtensions, err := apiextensionsclient.NewForConfig(apiConfig)
	if err != nil {
		panic(err)
	}

	return Kubernetes{
		logger:              logger,
		cfg:                 cfg,
		client:              clientset,
		clientApiExtensions: clientApiExtensions,
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

func (k *Kubernetes) GetNodes() []string {
	var ListEverything = v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
	res, err := k.client.CoreV1().Nodes().List(k.cfg.Ctx, ListEverything)
	if err != nil {
		panic(err.Error())
	}
	for _, s := range res.Items {
		k.logger.Info(s.Name)
	}
	return []string{}
}
