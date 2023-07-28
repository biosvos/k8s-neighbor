package domain

import "fmt"

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

func (g *GroupVersionKind) String() string {
	return fmt.Sprintf("%v.%v.%v", g.Group, g.Version, g.Kind)
}

type ResourceIdentifier struct {
	GVK       *GroupVersionKind
	Namespace string
	Name      string
}

func (r *ResourceIdentifier) String() string {
	return fmt.Sprintf("%v.%v.%v", r.GVK, r.Namespace, r.Name)
}
