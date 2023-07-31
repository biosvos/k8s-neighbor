package stakeholder

import (
	"github.com/biosvos/jason"
)

func getNamespace(contents []byte) string {
	namespaces := getStrings(contents, "metadata.namespace")
	if len(namespaces) == 0 {
		return ""
	}
	return namespaces[0]
}

func getStrings(contents []byte, path string) []string {
	root, err := jason.NewJason(contents)
	if err != nil {
		return nil
	}
	nodes := root.Path(path)
	//nodes, err := root.Parse(path)
	//if err != nil {
	//	return nil
	//}
	var ret []string
	for _, node := range nodes {
		want := node.String()
		ret = append(ret, want)
	}
	return ret
}
