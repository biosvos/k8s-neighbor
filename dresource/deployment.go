package dresource

import (
	"fmt"
)

var _ Resource = &Deployment{}

type Deployment struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
	Spec       struct {
		Selector *Selector `json:"selector"`
	} `json:"spec"`
}

func (d *Deployment) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", d.APIVersion, d.Kind, d.Metadata.Namespace, d.Metadata.Name)
}

func (d *Deployment) Relations() []*Relation {
	var ret []*Relation
	ret = append(ret, d.Metadata.Relations()...)
	ret = append(ret, NewPodSelectorRelation(d.Spec.Selector, d.Metadata.Namespace))
	return ret
}
