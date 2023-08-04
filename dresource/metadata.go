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
		ret = append(ret, NewNamespaceRelation(m.Namespace))
	}
	for _, reference := range m.OwnerReferences {
		ret = append(ret, NewOwnerReferenceRelation(reference, m.Namespace))
	}
	return ret
}
