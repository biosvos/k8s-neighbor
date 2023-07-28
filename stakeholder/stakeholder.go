package stakeholder

import "github.com/biosvos/k8s-neighbor/domain"

type Stakeholder interface {
	Find(contents []byte) ([]*domain.ResourceIdentifier, error)
}
