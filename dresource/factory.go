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
	case "apps/v1/Deployment", "apps/v1/ReplicaSet",
		"v1/Service",
		"apps/v1/DaemonSet", "apps/v1/StatefulSet",
		"batch/v1/Job", "batch/v1/CronJob":
		return NewResource[PodSpecSelector](contents)
	case "v1/Pod":
		return NewResource[Pod](contents)
	case "v1/ServiceAccount":
		return NewResource[ServiceAccount](contents)
	default:
		log.Println("unhandled resource", normal.Identity())
		return normal, nil
	}
}
