package dresource

import "strings"

type OwnerReference struct {
	APIVersion string `json:"apiVersion"`
	Controller bool   `json:"controller"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
}

func (o *OwnerReference) Group() string {
	if !strings.Contains(o.APIVersion, "/") {
		return ""
	}
	sp := strings.Split(o.APIVersion, "/")
	return sp[0]
}

func (o *OwnerReference) Version() string {
	if !strings.Contains(o.APIVersion, "/") {
		return o.APIVersion
	}
	sp := strings.Split(o.APIVersion, "/")
	return sp[1]
}

func (o *OwnerReference) Relation(namespace string) *Relation {
	return NewOwnerReferenceRelation(o, namespace)
}
