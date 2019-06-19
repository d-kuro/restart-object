package util

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Factory interface {
	ClientSet() (kubernetes.Interface, error)
}

type localFactoryImpl struct{}

func NewLocalFactory() Factory {
	return &localFactoryImpl{}
}

func (localFactoryImpl) ClientSet() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}

type inClusterFactoryImpl struct{}

func NewInClusterFactory() Factory {
	return &inClusterFactoryImpl{}
}

func (inClusterFactoryImpl) ClientSet() (kubernetes.Interface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
