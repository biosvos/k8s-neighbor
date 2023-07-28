package printer

import "github.com/biosvos/k8s-neighbor/domain"

type Printer interface {
	PrintResourceIdentifier(identifier *domain.ResourceIdentifier) error
	PrintResourceRelation(from *domain.ResourceIdentifier, to *domain.ResourceIdentifier) error
}
