package printer

import (
	"fmt"
	"github.com/biosvos/k8s-neighbor/dresource"
	"github.com/pkg/errors"
	"io"
)

var _ Printer = &D2{}

type D2 struct {
	Writer io.Writer
}

func (d *D2) PrintResourceIdentifier(resource dresource.Resource) error {
	_, err := io.WriteString(d.Writer, fmt.Sprintf("\"%v\"\n", resource.Identity()))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (d *D2) PrintResourceRelation(from dresource.Resource, to dresource.Resource, tag string) error {
	_, err := io.WriteString(d.Writer, fmt.Sprintf("\"%v\" -> \"%v\":%v\n", from.Identity(), to.Identity(), tag))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
