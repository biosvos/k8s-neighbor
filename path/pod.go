package path

var _ Finder = &Pod{}

type Pod struct{}

func (*Pod) ListPaths() []string {
	return []string{
		".metadata.namespace",
		".metadata.ownerReferences[*]", // ".apiVersion", ".name", ".kind"
		".spec.nodeName",
		".spec.serviceAccountName",
		".spec.volumes[*].configMap.name",
		".spec.volumes[*].projected.sources[*].configMap.name",
		".spec.serviceAccountName",
	}
}
