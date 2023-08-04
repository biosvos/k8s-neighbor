package dresource

type Metadata struct {
	Labels          map[string]string `json:"labels"`
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace"`
	OwnerReferences []*OwnerReference `json:"ownerReferences"`
}

func (m *Metadata) Relations() []*Relation {
	var ret []*Relation
	if m.Namespace != "" {
		ret = append(ret, &Relation{
			Type: NamespaceRelation,
			GVK: GroupVersionKind{
				Version: "v1",
				Kind:    "Namespace",
			},
			Name: m.Namespace,
		})
	}
	for _, reference := range m.OwnerReferences {
		ret = append(ret, &Relation{
			Type: OwnerReferenceRelation,
			GVK: GroupVersionKind{
				Group:   reference.Group(),
				Version: reference.Version(),
				Kind:    reference.Kind,
			},
			Namespace: m.Namespace,
			Name:      reference.Name,
		})
	}
	return ret
}
