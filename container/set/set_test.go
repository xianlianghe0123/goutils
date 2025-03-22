package set

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkSet(t *testing.T, s Set[int], expect []int) {
	require.Equal(t, len(expect), s.Size())
	for _, e := range expect {
		require.True(t, s.Contains(e))
	}
	cnt := 0
	for e := range s.Iter() {
		require.True(t, slices.Contains(expect, e))
		cnt++
	}
	require.Equal(t, len(expect), cnt)
}

// 测试 Set 类型的所有方法
func TestSet(t *testing.T) {
	s := NewSet[int]()
	checkSet(t, s, []int{})

	s.Add(1, 2, 3, 4)
	checkSet(t, s, []int{1, 2, 3, 4})

	require.True(t, s.Contains(1))
	require.False(t, s.Contains(0))

	s.Remove(1, 4)
	checkSet(t, s, []int{2, 3})

	// check iter break
	for i := range s.Iter() {
		if i == 2 {
			break
		}
	}

	s.Clear()
	checkSet(t, s, []int{})
}

func TestSets(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3, 4, 5, 6)
	other := NewSet[int]()
	other.Add(4, 5, 6, 7, 8)

	checkSet(t, s.Union(other), []int{1, 2, 3, 4, 5, 6, 7, 8})
	checkSet(t, s.Intersection(other), []int{4, 5, 6})
	checkSet(t, s.Difference(other), []int{1, 2, 3})
	checkSet(t, other.Difference(s), []int{7, 8})
}
