package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/stakeholder"
	"log"
)

func NewFlow(client *client.Client) *Flow {
	return &Flow{client: client}
}

type Flow struct {
	client *client.Client
}

func (f *Flow) GetWorkloadResources(group string, version string, kind string, namespace string, name string) {
	resource, err := f.client.Get(group, version, kind, namespace, name)
	if err != nil {
		panic(err)
	}
	stakeholders := stakeholder.MakeStakeholders(&domain.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	})
	for _, holder := range stakeholders {
		identifiers, err := holder.Find(resource)
		if err != nil {
			panic(err)
		}
		for _, identifier := range identifiers {
			log.Println(identifier)
		}
	}
}
