package json

import (
	"github.com/pkg/errors"
	"github.com/spyzhov/ajson"
)

func NewParser(contents []byte) (*Parser, error) {
	root, err := ajson.Unmarshal(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Parser{
		root: root,
	}, nil
}

type Parser struct {
	root *ajson.Node
}
