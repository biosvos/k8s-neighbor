package dresource

type RelationType uint64

const (
	UnknownRelation = iota
	SpecNameRelation
	OwnerReferenceRelation
	NamespaceRelation
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
