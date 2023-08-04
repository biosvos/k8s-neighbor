package client

import (
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

func (c *Client) Get(group, version, kind, namespace, name string) ([]byte, error) {
	uns := newUnstructured(group, version, kind, namespace, name)
	err := getResource(c.client, uns)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return jsonifyUnstructured(uns)
}

func (c *Client) ListByLabelSelector(group, version, kind, namespace string, selector map[string]string) ([][]byte, error) {
	unsList := newUnstructuredList(group, version, kind)
	err := listResourcesByLabelSelector(c.client, unsList, namespace, selector)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return jsonifyUnstructuredList(unsList)
}
