package query

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type DynamicKubeQuery struct {}


func NewDynamicQuery() DynamicKubeQuery {
	return DynamicKubeQuery{}
} 

func (dq DynamicKubeQuery) ListPods(kubeConfig *rest.Config, namespace string) error {

	dynamicClient, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		return err
	}

	resource := schema.GroupVersionResource{
		Version: "v1", 
		Resource: "pods",
	}

	pods, err := dynamicClient.Resource(resource).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	if len(pods.Items) == 0 {
		fmt.Printf("No Pods found in the %s namespace\n", namespace)
		return nil
	}

	fmt.Printf("%-40s%-20s%-20s\n", "NAME", "NAMESPACE", "STATUS")
	for _, pod := range pods.Items {
		
		// Pod specific
		name, found, err := unstructured.NestedString(pod.Object, "metadata", "name")
		if err != nil || !found {
			fmt.Printf("Not Found: %v\n", err)
			continue
		}

		status, found, err := unstructured.NestedString(pod.Object, "status", "phase")
		if err != nil || !found {
			fmt.Printf("Not Found: %v\n", err)
			continue
		}
		fmt.Printf("%-40s%-20s%-20s\n", name, namespace, status)
	}

	return nil
}