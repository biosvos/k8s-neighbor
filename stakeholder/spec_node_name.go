package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
)

type SpecNodeName struct{}

func (p *SpecNodeName) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	nodeNames := getStrings(contents, "spec.nodeName")
	var ret []*domain.ResourceIdentifier
	for _, name := range nodeNames {
		ret = append(ret, &domain.ResourceIdentifier{
			GVK: &domain.GroupVersionKind{
				Version: "v1",
				Kind:    "Node",
			},
			Name: name,
		})
	}
	return ret, nil
}
