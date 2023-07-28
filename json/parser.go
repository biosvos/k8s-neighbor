package json

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spyzhov/ajson"
	"strconv"
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

func (p *Parser) Parse(path string) ([]*Parser, error) {
	result, err := p.root.JSONPath(fmt.Sprintf("@%v", path))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var ret []*Parser
	for _, node := range result {
		ret = append(ret, &Parser{root: node})
	}
	return ret, nil
}

func (p *Parser) ParseOne(path string) (*Parser, error) {
	result, err := p.Parse(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(result) != 1 {
		return nil, errors.New("not one")
	}
	return result[0], nil
}

func (p *Parser) GetMustString() string {
	ret, err := p.GetString()
	if err != nil {
		panic(err)
	}
	return ret
}

func (p *Parser) GetString() (string, error) {
	switch p.root.Type() { //nolint:exhaustive
	case ajson.String:
		value, _ := p.root.GetString()
		return value, nil
	case ajson.Numeric:
		value, _ := p.root.GetNumeric()
		return strconv.FormatFloat(value, 'f', -1, 64), nil
	case ajson.Bool:
		value, _ := p.root.GetBool()
		return strconv.FormatBool(value), nil
	default:
		return "", errors.New("wrong type")
	}
}
