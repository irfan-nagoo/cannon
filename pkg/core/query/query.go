package query

import (
	"k8s.io/client-go/rest"
)


type Query interface {

	ListPods(kubeConfig *rest.Config, namespace string) error;
}

