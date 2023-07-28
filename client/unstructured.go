package client

import (
	"github.com/biosvos/k8s-neighbor/domain"
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

func jsonifyUnstructured(uns *unstructured.Unstructured) ([]byte, error) {
	json, err := uns.MarshalJSON()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return json, nil
}

func newUnstructuredList(identifier *domain.ResourceIdentifier) *unstructured.UnstructuredList {
	var ret unstructured.UnstructuredList
	ret.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   identifier.GVK.Group,
		Version: identifier.GVK.Version,
		Kind:    identifier.GVK.Kind,
	})
	return &ret
}

func jsonifyUnstructuredList(unsList *unstructured.UnstructuredList) ([][]byte, error) {
	var ret [][]byte
	for i := range unsList.Items {
		json, err := unsList.Items[i].MarshalJSON()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, json)
	}
	return ret, nil
}
