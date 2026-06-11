package query


import (
	"context"
	"fmt"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type KubeQuery struct {}

func NewQuery() DynamicKubeQuery {
	return DynamicKubeQuery{}
} 

func (kq KubeQuery) ListPods(kubeConfig *rest.Config, namespace string) error {

	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return err
	}

	pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	if len(pods.Items) == 0 {
		fmt.Printf("No Pods found in the %s namespace\n", namespace)
		return nil
	}

	fmt.Printf("%-40s%-20s%-20s\n", "NAME", "NAMESPACE", "STATUS")
	for _, pod := range pods.Items {
		fmt.Printf("%-40s%-20s%-20s\n", pod.Name, namespace, pod.Status.Phase)
	}

	return nil
}