package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/lib"
	"github.com/biosvos/k8s-neighbor/printer"
	"github.com/biosvos/k8s-neighbor/stakeholder"
	"log"
)

func NewFlow(client *client.Client, printer printer.Printer) *Flow {
	return &Flow{
		client:  client,
		printer: printer,
	}
}

type Flow struct {
	client  *client.Client
	printer printer.Printer
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

	conflict := map[string]struct{}{}
	for !stack.IsEmpty() {
		top := stack.Peek()
		stack.Drop(1)
		if _, ok := conflict[top.String()]; ok {
			continue
		}
		conflict[top.String()] = struct{}{}

		err := f.printer.PrintResourceIdentifier(top)
		if err != nil {
			panic(err)
		}

		resource, err := f.client.Get(top)
		if err != nil {
			log.Println(top, "not found")
			continue
		}
		stakeholders := stakeholder.MakeStakeholders(top.GVK)
		for _, holder := range stakeholders {
			identifiers, err := holder.Find(resource)
			if err != nil {
				panic(err)
			}
			for _, identifier := range identifiers {
				err := f.printer.PrintResourceRelation(top, identifier)
				if err != nil {
					panic(err)
				}
			}
			stack.Push(identifiers...)
		}
	}
}
