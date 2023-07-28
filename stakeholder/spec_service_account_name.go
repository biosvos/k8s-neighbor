package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
)

type SpecServiceAccountName struct{}

func (*SpecServiceAccountName) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	namespace := getNamespace(contents)
	serviceAccountNames := getStrings(contents, ".spec.serviceAccountName")
	var ret []*domain.ResourceIdentifier
	for _, name := range serviceAccountNames {
		ret = append(ret, &domain.ResourceIdentifier{
			GVK: &domain.GroupVersionKind{
				Version: "v1",
				Kind:    "ServiceAccount",
			},
			Namespace: namespace,
			Name:      name,
		})
	}
	return ret, nil
}
