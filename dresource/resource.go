package dresource

type Resource interface {
	Identity() string
	Relations() []*Relation
}
