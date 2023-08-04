package dresource

import (
	"fmt"
)

var _ Resource = &Normal{}

type Normal struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
}

func (n *Normal) Relations() []*Relation {
	return n.Metadata.Relations()
}

func (n *Normal) GVK() string {
	return fmt.Sprintf("%v/%v", n.APIVersion, n.Kind)
}

func (n *Normal) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", n.APIVersion, n.Kind, n.Metadata.Namespace, n.Metadata.Name)
}
