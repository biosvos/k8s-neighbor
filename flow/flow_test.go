package flow

import (
	"github.com/biosvos/k8s-neighbor/client"
	"testing"
)

func TestName(t *testing.T) {
	newClient, _ := client.NewClient()
	flow := NewFlow(newClient)
	flow.GetWorkloadResources("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")
}
