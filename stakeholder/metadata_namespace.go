package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/json"
	"github.com/pkg/errors"
)

type MetadataNamespace struct{}

func (m *MetadataNamespace) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	namespace := parseMetadataNamespace(parser)
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
