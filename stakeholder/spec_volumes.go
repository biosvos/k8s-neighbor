package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/json"
	"github.com/pkg/errors"
)

type SpecVolumes struct{}

func (*SpecVolumes) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	namespace := parseMetadataNamespace(parser)
	nodes := parseSpecVolumes(parser, namespace)
	return nodes, nil
}

func parseSpecVolumes(parser *json.Parser, namespace string) []*domain.ResourceIdentifier {
	volumeNodes, err := parser.Parse(".spec.volumes[*]")
	if err != nil {
		panic(err)
	}

	var ret []*domain.ResourceIdentifier
	for _, node := range volumeNodes {
		one, err := node.ParseOne(".configMap.name")
		if err != nil {
			continue
		}
		name := one.GetMustString()
		ret = append(ret, &domain.ResourceIdentifier{
			GVK: &domain.GroupVersionKind{
				Version: "v1",
				Kind:    "ConfigMap",
			},
			Namespace: namespace,
			Name:      name,
		})
	}
	return ret
}
