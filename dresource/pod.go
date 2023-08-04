package dresource

import "fmt"

var _ Resource = &Pod{}

type Pod struct {
	APIVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Metadata   *Metadata `json:"metadata"`
	Spec       struct {
		NodeName           string            `json:"nodeName"`
		NodeSelector       map[string]string `json:"nodeSelector"`
		PriorityClassName  string            `json:"priorityClassName"`
		SchedulerName      string            `json:"schedulerName"`
		ServiceAccountName string            `json:"serviceAccountName"`
		Volumes            []struct {
			ConfigMap *ConfigMap `json:"configMap,omitempty"`
			Projected struct {
				Sources []struct {
					ConfigMap *ConfigMap `json:"configMap,omitempty"`
				} `json:"sources,omitempty"`
			} `json:"projected,omitempty"`
		} `json:"volumes"`
	} `json:"spec"`
}

func (p *Pod) Identity() string {
	return fmt.Sprintf("%v/%v/%v/%v", p.APIVersion, p.Kind, p.Metadata.Namespace, p.Metadata.Name)
}

func (p *Pod) Relations() []*Relation {
	var ret []*Relation
	ret = append(ret, p.Metadata.Relations()...)
	ret = append(ret, NewNameSpecRelation("", "v1", "Node", p.Metadata.Namespace, p.Spec.NodeName))
	ret = append(ret, NewNameSpecRelation("", "v1", "ServiceAccount", p.Metadata.Namespace, p.Spec.ServiceAccountName))
	for _, volume := range p.Spec.Volumes {
		name := volume.ConfigMap.Name()
		if name != "" {
			ret = append(ret, NewNameSpecRelation("", "v1", "ConfigMap", p.Metadata.Namespace, name))
		}
		for _, source := range volume.Projected.Sources {
			name := source.ConfigMap.Name()
			if name != "" {
				ret = append(ret, NewNameSpecRelation("", "v1", "ConfigMap", p.Metadata.Namespace, name))
			}
		}
	}
	return ret
}
