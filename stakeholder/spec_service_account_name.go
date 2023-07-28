package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/json"
	"github.com/pkg/errors"
)

type SpecServiceAccountName struct{}

func (*SpecServiceAccountName) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	namespace := parseMetadataNamespace(parser)
	node := parseSpecServiceAccountName(parser, namespace)
	return []*domain.ResourceIdentifier{node}, nil
}

func parseSpecServiceAccountName(parser *json.Parser, namespace string) *domain.ResourceIdentifier {
	one, err := parser.ParseOne(".spec.serviceAccountName")
	if err != nil {
		panic(err)
	}
	name := one.GetMustString()
	return &domain.ResourceIdentifier{
		GVK: &domain.GroupVersionKind{
			Version: "v1",
			Kind:    "ServiceAccount",
		},
		Namespace: namespace,
		Name:      name,
	}
}
