package path

var _ Finder = &Deployment{}

type Deployment struct{}

func (*Deployment) ListPaths() []string {
	return []string{
		".metadata.namespace",
		".metadata.ownerReferences[*]", // ".apiVersion", ".name", ".kind"
		".spec.selector",
	}
}
