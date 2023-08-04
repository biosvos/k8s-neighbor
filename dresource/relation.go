package dresource

type RelationType string

const (
	UnknownRelation        RelationType = "UnknownRelation"
	SpecNameRelation                    = "SpecField"
	OwnerReferenceRelation              = "OwnerReference"
	NamespaceRelation                   = "Namespace"
	SelectorRelation                    = "Selector"
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
