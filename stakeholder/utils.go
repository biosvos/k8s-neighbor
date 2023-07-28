package stakeholder

import (
	"github.com/biosvos/k8s-neighbor/json"
)

func getNamespace(contents []byte) string {
	namespaces := getStrings(contents, ".metadata.namespace")
	if len(namespaces) == 0 {
		return ""
	}
	return namespaces[0]
}

func getStrings(contents []byte, path string) []string {
	parser, err := json.NewParser(contents)
	if err != nil {
		return nil
	}
	nodes, err := parser.Parse(path)
	if err != nil {
		return nil
	}
	var ret []string
	for _, node := range nodes {
		want, err := node.GetString()
		if err != nil {
			continue
		}
		ret = append(ret, want)
	}
	return ret
}
