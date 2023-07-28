package stakeholder

import "github.com/biosvos/k8s-neighbor/domain"

func MakeStakeholders(gvk *domain.GroupVersionKind) []Stakeholder {
	var ret []Stakeholder
	ret = append(ret,
		&MetadataOwnerReferences{},
		&MetadataNamespace{},
	)
	switch gvk.String() {
	case ".v1.Pod":
		ret = append(ret, &SpecNodeName{})
		ret = append(ret, &SpecServiceAccountName{})
	}
	return ret
}
