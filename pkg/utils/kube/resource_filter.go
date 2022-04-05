package kube

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type ResourceFilter interface {
	IsExcludedResource(group, kind, cluster string) bool
}

type ObjectFilter interface {
	IsExcludedObject(*unstructured.Unstructured) bool
}
