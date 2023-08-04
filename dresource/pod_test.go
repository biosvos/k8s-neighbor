package dresource

import (
	_ "embed"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed pod.json
var podJSON []byte

func TestPod(t *testing.T) {
	var pod Pod
	err := json.Unmarshal(podJSON, &pod)
	require.NoError(t, err)
	for _, relation := range pod.Relations() {
		t.Log(relation)
	}
}
