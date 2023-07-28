package stakeholder

import "github.com/biosvos/k8s-neighbor/json"

type Searcher interface {
	Search(parser *json.Parser) ([][]byte, error)
}
