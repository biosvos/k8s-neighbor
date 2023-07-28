package client

import (
	"context"
	"github.com/pkg/errors"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	clientGo "sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	client clientGo.Client
}

func (c *Client) Get(group string, version string, kind string, namespace string, name string) (string, error) {
	var uns unstructured.Unstructured
	uns.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	})
	uns.SetNamespace(namespace)
	uns.SetName(name)
	err := c.client.Get(context.Background(), clientGo.ObjectKeyFromObject(&uns), &uns)
	if err != nil {
		reason := apiErrors.ReasonForError(err)
		switch reason {
		case metaV1.StatusReasonNotFound:
			return "", ErrNotFound
		}
		return "", errors.WithStack(err)
	}
	json, err := uns.MarshalJSON()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(json), nil
}

func NewClient() (*Client, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	client, err := clientGo.New(config, clientGo.Options{})
	return &Client{
		client: client,
	}, nil
}
