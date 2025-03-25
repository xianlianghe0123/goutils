package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryIndexedTree(t *testing.T) {
	bit := NewBinaryIndexedTree([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(x int, y int) int { return x + y })
	require.Equal(t, []int{1, 3, 3, 10, 5, 11, 7, 36}, bit.tree)
	for i, r := range []int{1, 3, 6, 10, 15, 21, 28, 36} {
		require.Equal(t, r, bit.Accumulate(i))
	}
	bit.Renew(0, func(x int) int { return x - 1 })
	require.Equal(t, []int{0, 2, 3, 9, 5, 11, 7, 35}, bit.tree)
	require.Equal(t, []int{0, 2, 3, 4, 5, 6, 7, 8}, bit.origin)
	for i, r := range []int{0, 2, 5, 9, 14, 20, 27, 35} {
		require.Equal(t, r, bit.Accumulate(i))
	}
}

func TestBinaryIndexedTree_Range(t *testing.T) {
	origin := []int{8, 7, 6, 5, 4, 3, 2, 1}
	bit := NewBinaryIndexedTree(origin, func(x int, y int) int { return max(x, y) })
	for l := 0; l < len(origin); l++ {
		for r := l; r < len(origin); r++ {
			m := 0
			bit.Range(l, r, func(e int) { m = max(m, e) })
			require.Equal(t, origin[l], m)
		}
	}
}
