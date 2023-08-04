package dresource

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func NewResource[T any](contents []byte) (*T, error) {
	var ret T
	err := json.Unmarshal(contents, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ret, nil
}
