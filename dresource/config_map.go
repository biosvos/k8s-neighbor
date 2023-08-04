package dresource

type ConfigMap struct {
	ConfigMapName string `json:"name"`
}

func (c *ConfigMap) Name() string {
	if c == nil {
		return ""
	}
	return c.ConfigMapName
}
