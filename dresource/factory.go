package dresource

import (
	"github.com/pkg/errors"
	"log"
)

func MakeFactory(contents []byte) (Resource, error) {
	normal, err := NewResource[Normal](contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	switch normal.GVK() {
	case "apps/v1/Deployment":
		return NewResource[Deployment](contents)
	case "v1/Pod":
		return NewResource[Pod](contents)
	default:
		log.Println("unhandled resource", normal.GVK())
		return normal, nil
	}
}
