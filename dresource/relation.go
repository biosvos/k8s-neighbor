package dresource

type RelationType string

const (
	SpecNameRelation       RelationType = "SpecField"
	OwnerReferenceRelation RelationType = "OwnerReference"
	NamespaceRelation      RelationType = "Namespace"
	SelectorRelation       RelationType = "Selector"
)

type Relation struct {
	Type      RelationType
	GVK       GroupVersionKind
	Namespace string
	Name      string
	Selector  *Selector
}

func NewNameSpecRelation(group, version, kind, namespace, name string) *Relation {
	return &Relation{
		Type: SpecNameRelation,
		GVK: GroupVersionKind{
			Group:   group,
			Version: version,
			Kind:    kind,
		},
		Namespace: namespace,
		Name:      name,
	}
}

func NewPodSelectorRelation(selector *Selector, namespace string) *Relation {
	return &Relation{
		Type: SelectorRelation,
		GVK: GroupVersionKind{
			Version: "v1",
			Kind:    "Pod",
		},
		Namespace: namespace,
		Selector:  selector,
	}
}

func NewNamespaceRelation(namespace string) *Relation {
	return &Relation{
		Type: NamespaceRelation,
		GVK: GroupVersionKind{
			Version: "v1",
			Kind:    "Namespace",
		},
		Name: namespace,
	}
}

func NewOwnerReferenceRelation(reference *OwnerReference, namespace string) *Relation {
	return &Relation{
		Type: OwnerReferenceRelation,
		GVK: GroupVersionKind{
			Group:   reference.Group(),
			Version: reference.Version(),
			Kind:    reference.Kind,
		},
		Namespace: namespace,
		Name:      reference.Name,
	}
}
