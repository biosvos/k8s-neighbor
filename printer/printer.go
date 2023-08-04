package printer

import (
	"github.com/biosvos/k8s-neighbor/dresource"
)

type Printer interface {
	PrintResourceIdentifier(resource dresource.Resource) error
	PrintResourceRelation(from dresource.Resource, to dresource.Resource) error
}
