package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()
	for _, u := range [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 0}, {2, 4}, {1, 3}, {5, 7}} {
		uf.Union(u[0], u[1])
	}

	for _, c := range [][]int{
		{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
		{5, 6}, {5, 7}, {5, 8}, {6, 7}, {6, 8}, {7, 8},
		{9, 0},
	} {
		require.True(t, uf.IsConnect(c[0], c[1]))
	}

	for _, c := range [][]int{
		{1, 5}, {1, 7}, {1, 0},
		{2, 6}, {2, 8},
		{9, 4}, {0, 6},
	} {
		require.False(t, uf.IsConnect(c[0], c[1]))
	}

	require.Equal(t, 3, uf.ConnectedComponent())
}
