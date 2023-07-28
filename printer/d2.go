package printer

import (
	"fmt"
	"github.com/biosvos/k8s-neighbor/domain"
	"github.com/pkg/errors"
	"io"
)

var _ Printer = &D2{}

type D2 struct {
	Writer io.Writer
}

func (d *D2) PrintResourceIdentifier(identifier *domain.ResourceIdentifier) error {
	_, err := io.WriteString(d.Writer, fmt.Sprintf("\"%v\"\n", identifier))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (d *D2) PrintResourceRelation(from *domain.ResourceIdentifier, to *domain.ResourceIdentifier) error {
	_, err := io.WriteString(d.Writer, fmt.Sprintf("\"%v\" -> \"%v\"\n", from, to))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
