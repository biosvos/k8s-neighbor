package dresource

import "fmt"

var _ Resource = &ServiceAccount{}

type ServiceAccount struct {
	Kind       string    `json:"kind"`
	APIVersion string    `json:"apiVersion"`
	Metadata   *Metadata `json:"metadata"`
	Secrets    []struct {
		Name string `json:"name"`
	} `json:"secrets,omitempty"`
	ImagePullSecrets []struct {
		Name string `json:"name"`
	} `json:"imagePullSecrets,omitempty"`
}

func (s *ServiceAccount) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", s.APIVersion, s.Kind, s.Metadata.Namespace, s.Metadata.Name)
}

func (s *ServiceAccount) Relations() []*Relation {
	var ret []*Relation
	ret = append(ret, s.Metadata.Relations()...)
	for _, secret := range s.Secrets {
		ret = append(ret, NewNameSpecRelation("", "v1", "Secret", s.Metadata.Namespace, secret.Name))
	}
	for _, secret := range s.ImagePullSecrets {
		ret = append(ret, NewNameSpecRelation("", "v1", "Secret", s.Metadata.Namespace, secret.Name))
	}
	return ret
}
