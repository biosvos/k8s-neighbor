package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewParser(t *testing.T) {
	parser, err := NewParser([]byte(`{
  "metadata": {
    "ownerReferences": [
      {
        "apiVersion": "apps/v1",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "ReplicaSet",
        "name": "coredns-59b4f5bbd5",
        "uid": "99e999e3-9f2d-4d35-8aac-9abd9bd6928a"
      }
    ]
  }
}`))
	require.NoError(t, err)
	require.NotNil(t, parser)
}

func TestParse(t *testing.T) {
	parser, _ := NewParser([]byte(`{
  "metadata": {
    "ownerReferences": [
      {
        "apiVersion": "apps/v1",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "ReplicaSet",
        "name": "coredns-59b4f5bbd5",
        "uid": "99e999e3-9f2d-4d35-8aac-9abd9bd6928a"
      }
    ]
  }
}`))

	nodes, err := parser.Parse(".metadata.ownerReferences[*]")

	require.NoError(t, err)
	require.Len(t, nodes, 1)

	for _, node := range nodes {
		kinds, _ := node.Parse(".kind")
		getString, err := kinds[0].GetString()

		require.NoError(t, err)
		require.Equal(t, "ReplicaSet", getString)

		blockOwnerDeletion, _ := node.ParseOne(".blockOwnerDeletion")
		getString, err = blockOwnerDeletion.GetString()

		require.NoError(t, err)
		require.Equal(t, "true", getString)
	}
}
