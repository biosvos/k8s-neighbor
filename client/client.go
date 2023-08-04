package client

import (
	"github.com/biosvos/k8s-neighbor/dresource"
	"github.com/pkg/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	ctrl "sigs.k8s.io/controller-runtime"
	clientGo "sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
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
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		client:    client,
		discovery: discoveryClient,
	}, nil
}

type Client struct {
	client    clientGo.Client
	discovery *discovery.DiscoveryClient
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

func (c *Client) List() ([][]byte, error) {
	gvks, err := c.listGVKs()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret [][]byte
	for _, gvk := range gvks {
		unsList := newUnstructuredList(gvk.Group, gvk.Version, gvk.Kind)
		err := listResources(c.client, unsList)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		list, err := jsonifyUnstructuredList(unsList)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, list...)
	}
	return ret, nil
}

func (c *Client) listGVKs() ([]*dresource.GroupVersionKind, error) {
	resources, err := c.discovery.ServerPreferredResources()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret []*dresource.GroupVersionKind
	for _, resource := range resources {
		group, version := parseAPIVersion(resource.GroupVersion)
		for _, apiResource := range resource.APIResources {
			if !hasListVerb(apiResource.Verbs) {
				continue
			}
			ret = append(ret, &dresource.GroupVersionKind{
				Group:   group,
				Version: version,
				Kind:    apiResource.Kind,
			})
		}
	}
	return ret, nil
}

func hasListVerb(verbs metaV1.Verbs) bool {
	for _, verb := range verbs {
		if verb == "list" {
			return true
		}
	}
	return false
}

func parseAPIVersion(apiVersion string) (string, string) {
	split := strings.Split(apiVersion, "/")
	if len(split) == 1 {
		return "", split[0]
	}
	return split[0], split[1]
}
