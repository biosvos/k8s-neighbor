package client

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, _ := NewClient()

		resource, err := client.Get("", "v1", "Pod", "kube-system", "coredns-59b4f5bbd5-kbrvk")

		require.NoError(t, err)
		require.NotEmpty(t, resource)
	})

	t.Run("unknown namespace", func(t *testing.T) {
		client, _ := NewClient()

		resource, err := client.Get("", "v1", "Pod", "kak", "coredns-59b4f5bbd5-kbrvk")

		require.ErrorIs(t, err, ErrNotFound)
		require.Empty(t, resource)
	})
}

func TestAll(t *testing.T) {
	client, _ := NewClient()
	resources, err := client.List()
	require.NoError(t, err)
	require.NotNil(t, resources)
}
