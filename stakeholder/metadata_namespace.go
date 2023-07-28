package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
)

type MetadataNamespace struct{}

func (m *MetadataNamespace) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	namespace := getNamespace(contents)
	if namespace == "" {
		return nil, nil
	}
	return []*domain.ResourceIdentifier{
		{
			GVK: &domain.GroupVersionKind{
				Group:   "",
				Version: "v1",
				Kind:    "Namespace",
			},
			Name: namespace,
		},
	}, nil
}
