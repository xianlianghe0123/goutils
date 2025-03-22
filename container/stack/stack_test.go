package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkStack(t *testing.T, s *Stack[int], expect []int) {
	require.Equal(t, len(expect), s.Size())
	require.Equal(t, len(expect) == 0, s.IsEmpty())
	require.Equal(t, expect, s.elems)
	if s.Size() > 0 {
		require.Equal(t, expect[len(expect)-1], s.Top())
	}
	i := len(expect) - 1
	for e := range s.Iter() {
		require.Equal(t, expect[i], e)
		i--
	}
}

func TestStack(t *testing.T) {
	s := NewStack[int](10)
	require.Equal(t, 10, cap(s.elems))
	checkStack(t, s, []int{})

	s.Push(1, 2)
	checkStack(t, s, []int{1, 2})

	require.Equal(t, 2, s.Pop())
	checkStack(t, s, []int{1})

	require.Equal(t, 1, s.Pop())
	checkStack(t, s, []int{})

	s.Push(3, 4, 5)
	checkStack(t, s, []int{3, 4, 5})
	for e := range s.Iter() {
		if e == 4 {
			break
		}
	}

	s.Clear()
	checkStack(t, s, []int{})

	require.Panics(t, func() { s.Pop() })
	require.Panics(t, func() { s.Top() })
}
