package client

import (
	"context"
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
	switch reason {
	case metaV1.StatusReasonNotFound:
		return ErrNotFound
	default:
		return err
	}
}
