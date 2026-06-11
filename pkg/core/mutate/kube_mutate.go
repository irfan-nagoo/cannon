package mutate

import (
	
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
)


func SetNamespace(namespace string) error {

	var kubeConfigPath string
	if homeDir := os.Getenv("HOME"); homeDir != "" {
		kubeConfigPath = filepath.Join(homeDir, ".kube", "config")
	}

	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return err
	}
	previousNamespace := config.Contexts[config.CurrentContext].Namespace
	config.Contexts[config.CurrentContext].Namespace = namespace
	clientcmd.WriteToFile(*config, kubeConfigPath)
	
	fmt.Printf("Namesapce switched from %s to %s\n", previousNamespace, namespace)

	return  nil
}