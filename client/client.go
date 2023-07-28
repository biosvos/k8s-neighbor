package client

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	clientGo "sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient() (*Client, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	client, err := clientGo.New(config, clientGo.Options{})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		client: client,
	}, nil
}

type Client struct {
	client clientGo.Client
}

func (c *Client) Get(identifier *domain.ResourceIdentifier) ([]byte, error) {
	uns := newUnstructured(identifier.GVK.Group, identifier.GVK.Version, identifier.GVK.Kind, identifier.Namespace, identifier.Name)
	err := getResource(c.client, uns)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return jsonifyUnstructured(uns)
}
