package path

import "github.com/biosvos/k8s-neighbor/json"

type Path interface {
	Source() string
	Structure() any
}

type realPath struct {
	path    string
	element any
}

func (r *realPath) Source() string {
	return r.path
}

func (r *realPath) Structure() any {
	return r.element
}

func Work(path Path, contents []byte) {
	parser, err := json.NewParser(contents)
	if err != nil {
		return
	}
	parse, err := parser.Parse(path.Source())
	if err != nil {
		return
	}
}

type Finder interface {
	ListPaths() []string
}
