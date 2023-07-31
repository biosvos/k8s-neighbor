package stakeholder

import (
	"github.com/biosvos/jason"
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type MetadataOwnerReferences struct{}

func (o *MetadataOwnerReferences) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	root, err := jason.NewJason(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	namespace := getNamespace(contents)
	ret := parseOwnerReferences(root, namespace)
	return ret, nil
}

func parseOwnerReferences(root jason.Node, namespace string) []*domain.ResourceIdentifier {
	nodes := root.Path("metadata.ownerReferences.*")
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

func parseAPIVersion(node jason.Node) (string, string) {
	apiVersion, err := node.Get("apiVersion")
	if err != nil {
		panic(err)
	}
	gv, err := schema.ParseGroupVersion(apiVersion.String())
	if err != nil {
		panic(err)
	}
	return gv.Group, gv.Version
}

func parseName(node jason.Node) string {
	nameNode, err := node.Get("name")
	if err != nil {
		panic(err)
	}
	return nameNode.String()
}

func parseKind(node jason.Node) string {
	kindNode, err := node.Get("kind")
	if err != nil {
		panic(err)
	}
	return kindNode.String()
}
