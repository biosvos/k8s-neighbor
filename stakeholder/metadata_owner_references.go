package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/json"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type MetadataOwnerReferences struct{}

func (o *MetadataOwnerReferences) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	namespace := parseMetadataNamespace(parser)
	return parseOwnerReferences(parser, namespace), nil
}

func parseMetadataNamespace(node *json.Parser) string {
	one, err := node.ParseOne(".metadata.namespace")
	if err != nil {
		return ""
	}
	return one.GetMustString()
}

func parseOwnerReferences(parser *json.Parser, namespace string) []*domain.ResourceIdentifier {
	nodes, err := parser.Parse(".metadata.ownerReferences[*]")
	if err != nil {
		panic(err)
	}
	var ret []*domain.ResourceIdentifier
	for _, node := range nodes {
		group, version := parseAPIVersion(node)
		kind := parseKind(node)
		name := parseName(node)
		ret = append(ret, &domain.ResourceIdentifier{
			GVK: &domain.GroupVersionKind{
				Group:   group,
				Version: version,
				Kind:    kind,
			},
			Namespace: namespace,
			Name:      name,
		})
	}
	return ret
}

func parseAPIVersion(node *json.Parser) (string, string) {
	apiVersion, err := node.ParseOne(".apiVersion")
	if err != nil {
		panic(err)
	}
	gv, err := schema.ParseGroupVersion(apiVersion.GetMustString())
	if err != nil {
		panic(err)
	}
	return gv.Group, gv.Version
}

func parseName(node *json.Parser) string {
	nameNode, err := node.ParseOne(".name")
	if err != nil {
		panic(err)
	}
	return nameNode.GetMustString()
}

func parseKind(node *json.Parser) string {
	kindNode, err := node.ParseOne(".kind")
	if err != nil {
		panic(err)
	}
	return kindNode.GetMustString()
}
