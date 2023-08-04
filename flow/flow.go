package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/biosvos/k8s-neighbor/dresource"
	"github.com/biosvos/k8s-neighbor/lib"
	"github.com/biosvos/k8s-neighbor/printer"
	"github.com/pkg/errors"
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

type Conflict struct {
	elements map[string]struct{}
}

func NewConflict() *Conflict {
	return &Conflict{
		elements: map[string]struct{}{},
	}
}

func (c *Conflict) Check(resource dresource.Resource) bool {
	key := resource.Identity()
	_, ok := c.elements[key]
	c.elements[key] = struct{}{}
	return ok
}

type Pair struct {
	FromResource dresource.Resource
	ToRelation   *dresource.Relation
}

func (f *Flow) GetWorkloadResources(group string, version string, kind string, namespace string, name string) {
	stack := lib.NewStack[*Pair]()
	stack.Push(&Pair{
		FromResource: nil,
		ToRelation:   dresource.NewNameSpecRelation(group, version, kind, namespace, name),
	})

	conflict := NewConflict()
	for !stack.IsEmpty() {
		top := stack.Peek()
		stack.Drop(1)

		resources, err := Query(f.client, top.ToRelation)
		if err != nil {
			if errors.Is(err, client.ErrNotFound) {
				continue
			}
			panic(err)
		}

		for _, resource := range resources {
			f.print(top, resource)
			if conflict.Check(resource) {
				continue
			}

			relations := resource.Relations()
			for _, relation := range relations {
				stack.Push(&Pair{
					FromResource: resource,
					ToRelation:   relation,
				})
			}
		}
	}
}

func (f *Flow) print(top *Pair, resource dresource.Resource) {
	err := f.printer.PrintResourceIdentifier(resource)
	if err != nil {
		panic(err)
	}
	if top.FromResource == nil {
		return
	}
	err = f.printer.PrintResourceRelation(top.FromResource, resource, string(top.ToRelation.Type))
	if err != nil {
		panic(err)
	}
}

func Query(c *client.Client, top *dresource.Relation) ([]dresource.Resource, error) {
	switch top.Type {
	case dresource.SpecNameRelation, dresource.NamespaceRelation, dresource.OwnerReferenceRelation:
		content, err := c.Get(top.GVK.Group, top.GVK.Version, top.GVK.Kind, top.Namespace, top.Name)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		resource, err := dresource.MakeFactory(content)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return []dresource.Resource{resource}, nil
	case dresource.SelectorRelation:
		contents, err := c.ListByLabelSelector(top.GVK.Group, top.GVK.Version, top.GVK.Kind, top.Namespace, top.Selector.MatchLabels)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		var ret []dresource.Resource
		for _, content := range contents {
			resource, err := dresource.MakeFactory(content)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			ret = append(ret, resource)
		}
		return ret, nil
	default:
		panic("sdf")
	}
}
