package stakeholder

import "github.com/biosvos/k8s-neighbor/domain"

func MakeStakeholders(gvk *domain.GroupVersionKind) []Stakeholder {
	var ret []Stakeholder
	ret = append(ret, &OwnerReferences{})
	switch gvk.String() {
	case ".v1.Pod":
		ret = append(ret, &PodNodeName{})
	}
	return ret
}
