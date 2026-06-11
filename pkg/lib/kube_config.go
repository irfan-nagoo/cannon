package lib

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func GetKubeConfig() (*rest.Config, error) {
	
	var kubeConfigPath string
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		kubeConfigPath = filepath.Join(homeDir, ".kube", "config")
	}
	
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	kubeConfig.QPS = 50
	kubeConfig.Burst = 100
	
	return kubeConfig, nil
}

func GetCurrentNamespace() (string, error) {

	var kubeConfigPath string
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		kubeConfigPath = filepath.Join(homeDir, ".kube", "config")
	}

	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return "", err
	}
	
	return config.Contexts[config.CurrentContext].Namespace, nil
}