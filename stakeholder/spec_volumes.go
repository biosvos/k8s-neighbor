package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
)

type SpecVolumes struct{}

func (*SpecVolumes) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	namespace := getNamespace(contents)
	var configMapNames []string
	strings := getStrings(contents, "spec.volumes[*].configMap.name")
	configMapNames = append(configMapNames, strings...)
	strings = getStrings(contents, "spec.volumes[*].projected.sources[*].configMap.name")
	configMapNames = append(configMapNames, strings...)
	var ret []*domain.ResourceIdentifier
	for _, name := range configMapNames {
		ret = append(ret, &domain.ResourceIdentifier{
			GVK: &domain.GroupVersionKind{
				Version: "v1",
				Kind:    "ConfigMap",
			},
			Namespace: namespace,
			Name:      name,
		})
	}
	return ret, nil
}
