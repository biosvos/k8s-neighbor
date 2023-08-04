package dresource

import "fmt"

var _ Resource = &RoleBinding{}

type RoleBinding struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
	RoleRef    struct {
		APIGroup APIGroup `json:"apiGroup"`
		Kind     string   `json:"kind"`
		Name     string   `json:"name"`
	} `json:"roleRef"`
	Subjects []struct {
		APIGroup  APIGroup `json:"apiGroup"`
		Kind      string   `json:"kind"`
		Name      string   `json:"name"`
		Namespace string   `json:"namespace"`
	} `json:"subjects"`
}

func (r *RoleBinding) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", r.APIVersion, r.Kind, r.Metadata.Namespace, r.Metadata.Name)
}

func (r *RoleBinding) Relations() []*Relation {
	var ret []*Relation
	ret = append(ret, r.Metadata.Relations()...)
	group := r.RoleRef.APIGroup.Group()
	if group == "" { // default value
		group = "rbac.authorization.k8s.io"
	}
	ret = append(ret, NewNameSpecRelation(group, r.RoleRef.APIGroup.Version(), r.RoleRef.Kind, r.Metadata.Namespace, r.RoleRef.Name))
	for _, subject := range r.Subjects {
		ret = append(ret, NewNameSpecRelation(subject.APIGroup.Group(), subject.APIGroup.Version(), subject.Kind, subject.Namespace, subject.Name))
	}
	return ret
}
