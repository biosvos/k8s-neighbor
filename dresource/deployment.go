package dresource

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

var _ Resource = &Deployment{}

func NewDeployment(contents []byte) (*Deployment, error) {
	var ret Deployment
	err := json.Unmarshal(contents, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ret, nil
}

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
	ret = append(ret, &Relation{
		Type: SelectorRelation,
		GVK: GroupVersionKind{
			Version: "v1",
			Kind:    "Pod",
		},
		Namespace: d.Metadata.Namespace,
		//Name:      "",
		Selector: d.Spec.Selector,
	})
	return ret
}
