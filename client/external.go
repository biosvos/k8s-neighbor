package client

import (
	"context"
	"github.com/pkg/errors"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	clientGo "sigs.k8s.io/controller-runtime/pkg/client"
)

func getResource(client clientGo.Client, uns *unstructured.Unstructured) error {
	err := client.Get(context.Background(), clientGo.ObjectKeyFromObject(uns), uns)
	if err != nil {
		return mappingError(err)
	}
	return nil
}

func mappingError(err error) error {
	reason := apiErrors.ReasonForError(err)
	switch reason { //nolint:exhaustive
	case metaV1.StatusReasonNotFound:
		return ErrNotFound
	default:
		return err
	}
}

func listResourcesByLabelSelector(client clientGo.Client, unsList *unstructured.UnstructuredList, namespace string, selector map[string]string) error {
	err := client.List(context.Background(), unsList, clientGo.InNamespace(namespace), clientGo.MatchingLabels(selector))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func listResources(client clientGo.Client, unsList *unstructured.UnstructuredList) error {
	err := client.List(context.Background(), unsList)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
