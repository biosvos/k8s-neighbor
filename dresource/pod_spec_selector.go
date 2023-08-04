package dresource

import (
	"fmt"
)

var _ Resource = &PodSpecSelector{}

type PodSpecSelector struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
	Spec       struct {
		Selector *Selector `json:"selector"`
	} `json:"spec"`
}

func (p *PodSpecSelector) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", p.APIVersion, p.Kind, p.Metadata.Namespace, p.Metadata.Name)
}

func (p *PodSpecSelector) Relations() []*Relation {
	var ret []*Relation
	ret = append(ret, p.Metadata.Relations()...)
	ret = append(ret, NewPodSelectorRelation(p.Spec.Selector, p.Metadata.Namespace))
	return ret
}
