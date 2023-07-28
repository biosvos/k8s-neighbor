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
	return &Client{
		client: client,
	}, nil
}

type Client struct {
	client clientGo.Client
}

func (c *Client) Get(group string, version string, kind string, namespace string, name string) (string, error) {
	uns := newUnstructured(group, version, kind, namespace, name)
	err := getResource(c.client, uns)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return jsonifyUnstructured(uns)
}
