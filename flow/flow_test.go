package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	newClient, _ := client.NewClient()
	flow := NewFlow(newClient)
	require.NotNil(t, flow)
	flow.GetWorkloadResources("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")
}
