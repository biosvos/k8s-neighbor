package client

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, client)

	resource, err := client.Get("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")

	require.NoError(t, err)
	require.NotEmpty(t, resource)
}
