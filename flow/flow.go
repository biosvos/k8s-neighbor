package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/lib"
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
	stack := lib.NewStack[*domain.ResourceIdentifier]()
	stack.Push(&domain.ResourceIdentifier{
		GVK: &domain.GroupVersionKind{
			Group:   group,
			Version: version,
			Kind:    kind,
		},
		Namespace: namespace,
		Name:      name,
	})

	for !stack.IsEmpty() {
		top := stack.Peek()
		log.Println("current:", top)
		stack.Drop(1)

		resource, err := f.client.Get(top)
		if err != nil {
			panic(err)
		}
		stakeholders := stakeholder.MakeStakeholders(top.GVK)
		for _, holder := range stakeholders {
			identifiers, err := holder.Find(resource)
			if err != nil {
				panic(err)
			}
			for _, identifier := range identifiers {
				log.Println(identifier)
			}
			stack.Push(identifiers...)
		}
	}
}
