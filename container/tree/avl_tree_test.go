package tree

import (
	"cmp"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkAvlTree(t *testing.T, avl *AvlTree[int, int], expect []int, msg ...any) {
	require.Equal(t, len(expect), avl.Size(), msg...)
	for i := range expect {
		require.True(t, avl.Contains(expect[i]), msg...)
		e, ok := avl.Get(expect[i])
		require.True(t, ok, msg...)
		require.Equal(t, e, expect[i], msg...)
		e, ok = avl.Rank(i + 1)
		require.True(t, ok, msg...)
		require.Equal(t, expect[i], e, msg...)
	}
	lastK, lastV := math.MinInt32, math.MinInt32
	for k, v := range avl.Iter() {
		require.True(t, k > lastK, msg...)
		require.True(t, v > lastV, msg...)
		lastK, lastV = k, v
	}
}

func TestAvlTree_ReadNotExist(t *testing.T) {
	avl := NewAvlTree[int, int](cmp.Compare)

	require.False(t, avl.Contains(1))
	_, ok := avl.Get(1)
	require.False(t, ok)
	_, ok = avl.Rank(1)
	require.False(t, ok)

	avl.Set(1, 1)
	require.False(t, avl.Contains(2))
	_, ok = avl.Get(2)
	require.False(t, ok)
	_, ok = avl.Rank(2)
	require.False(t, ok)
}

func buildAvlTree(elems []int) *AvlTree[int, int] {
	avl := NewAvlTree[int, int](cmp.Compare)
	for _, elem := range elems {
		avl.Set(elem, elem)
	}
	return avl
}

func TestAvlTree_Set(t *testing.T) {
	// empty
	avl := NewAvlTree[int, int](cmp.Compare)
	avl.Set(1, 1)
	checkAvlTree(t, avl, []int{1})
	// same
	avl = buildAvlTree([]int{2, 1, 3})
	avl.Set(2, 2)
	checkAvlTree(t, avl, []int{1, 2, 3})
	// ll
	//        |           |
	//        4           2
	//       / \         / \
	//      2   5  =>   1   4
	//     / \         /   / \
	//    1   3       0   3   5
	//   /
	//  0
	avl = buildAvlTree([]int{4, 5, 2, 1, 3})
	avl.Set(0, 0)
	checkAvlTree(t, avl, []int{0, 1, 2, 3, 4, 5})
	// rr
	//     |             |
	//     1             3
	//    / \           / \
	//   0   3    =>   1   4
	//      / \       / \   \
	//     2   4     0   2   5
	//          \
	//           5
	avl = buildAvlTree([]int{1, 0, 3, 2, 4})
	avl.Set(5, 5)
	checkAvlTree(t, avl, []int{0, 1, 2, 3, 4, 5})
	// lr
	//      |            |          |
	//      4            4          3
	//     / \          / \        / \
	//    1   5   =>   3   5  =>  1   4
	//   / \          /          / \   \
	//  0   3        1          0   2   5
	//     /        / \
	//    2        0   2
	avl = buildAvlTree([]int{4, 5, 1, 0, 3})
	avl.Set(2, 2)
	checkAvlTree(t, avl, []int{0, 1, 2, 3, 4, 5})
	// rl
	//     | 	        |            |
	//     1            1            2
	//    / \          / \          / \
	//   0   4   =>   0   2   =>   1   4
	//      / \            \      /   / \
	//     2   5            4    0   3   5
	//      \              / \
	//       3            3   5
	avl = buildAvlTree([]int{1, 4, 0, 2, 5})
	avl.Set(3, 3)
	checkAvlTree(t, avl, []int{0, 1, 2, 3, 4, 5})
}

func TestAvlTree_Remove(t *testing.T) {
	// empty
	avl := NewAvlTree[int, int](cmp.Compare)
	avl.Remove(1)
	checkAvlTree(t, avl, []int{})
	// not exist
	avl = buildAvlTree([]int{2, 1, 3})
	avl.Remove(4)
	checkAvlTree(t, avl, []int{1, 2, 3})
	// leaf
	elems := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13}
	for i := range elems {
		avl = buildAvlTree(elems)
		avl.Remove(i + 1)
		expect := make([]int, 0, 16)
		for j := range elems {
			if j != i {
				expect = append(expect, j+1)
			}
		}
		checkAvlTree(t, avl, expect, "case", i)
	}
}

func TestAvlTree_IterBreak(t *testing.T) {
	avl := buildAvlTree([]int{2, 1, 3})
	for i := range 3 {
		for e := range avl.Iter() {
			if e == i+1 {
				break
			}
		}
	}
}
