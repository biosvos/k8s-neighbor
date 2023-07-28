package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/biosvos/k8s-neighbor/json"
	"github.com/pkg/errors"
)

type PodNodeName struct{}

func (p *PodNodeName) Find(contents []byte) ([]*domain.ResourceIdentifier, error) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	node := parseNode(parser)
	return []*domain.ResourceIdentifier{node}, nil
}

func parseNode(parser *json.Parser) *domain.ResourceIdentifier {
	one, err := parser.ParseOne(".spec.nodeName")
	if err != nil {
		panic(err)
	}
	name := one.GetMustString()
	return &domain.ResourceIdentifier{
		GVK: &domain.GroupVersionKind{
			Version: "v1",
			Kind:    "Node",
		},
		Name: name,
	}
}
