package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkQueue(t *testing.T, q *Queue[int], expect []int) {
	require.Equal(t, len(expect), q.Size())
	require.Equal(t, len(expect) == 0, q.IsEmpty())
	if len(expect) > 0 {
		require.Equal(t, expect[0], q.Front())
	}
	i := 0
	for e := range q.Iter() {
		require.Equal(t, expect[i], e)
		i++
	}
}

func TestQueue(t *testing.T) {
	_BlockCapability = 4
	q := NewQueue[int]()
	checkQueue(t, q, []int{})

	q.Push(1, 2, 3, 4, 5)
	checkQueue(t, q, []int{1, 2, 3, 4, 5})

	q.Push(6, 7, 8, 9, 10)
	checkQueue(t, q, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	require.Equal(t, 1, q.Pop())
	require.Equal(t, 2, q.Pop())
	checkQueue(t, q, []int{3, 4, 5, 6, 7, 8, 9, 10})

	require.Equal(t, 3, q.Pop())
	require.Equal(t, 4, q.Pop())
	require.Equal(t, 5, q.Pop())
	require.Equal(t, 6, q.Pop())
	checkQueue(t, q, []int{7, 8, 9, 10})

	// check iter break
	for e := range q.Iter() {
		if e == 8 {
			break
		}
	}

	q.Clear()
	checkQueue(t, q, []int{})

	require.Panics(t, func() { q.Front() })
	require.Panics(t, func() { q.Pop() })
}
