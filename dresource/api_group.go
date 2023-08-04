package dresource

import "strings"

type APIGroup string

func (a APIGroup) Group() string {
	if !strings.Contains(string(a), "/") {
		return ""
	}
	sp := strings.Split(string(a), "/")
	return sp[0]
}

func (a APIGroup) Version() string {
	if !strings.Contains(string(a), "/") {
		return string(a)
	}
	sp := strings.Split(string(a), "/")
	return sp[1]
}

func parseAPIVersion(apiVersion string) (string, string) {
	split := strings.Split(apiVersion, "/")
	if len(split) == 1 {
		return "", split[0]
	}
	return split[0], split[1]
}
