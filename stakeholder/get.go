package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/json"
)

var _ Searcher = &Get{}

type Get struct {
}

func (g *Get) Search(parser *json.Parser) ([][]byte, error) {
	//TODO implement me
	panic("implement me")
}
