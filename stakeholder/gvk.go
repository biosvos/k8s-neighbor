package stakeholder

type Query struct {
	Path     string
	Searcher Searcher
}

type Resource interface {
	ListQueries() []*Query
}

var _ Resource = &Pod{}

type Pod struct{}

func (p *Pod) ListQueries() []*Query {
	return []*Query{
		{
			Path:     ".metadata.namespace",
			Searcher: nil,
		},
	}
}
