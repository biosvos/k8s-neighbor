package client

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func newUnstructured(group string, version string, kind string, namespace string, name string) *unstructured.Unstructured {
	var ret unstructured.Unstructured
	ret.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	})
	ret.SetNamespace(namespace)
	ret.SetName(name)
	return &ret
}

func jsonifyUnstructured(uns *unstructured.Unstructured) (string, error) {
	json, err := uns.MarshalJSON()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(json), nil
}
