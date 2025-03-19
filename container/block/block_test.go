package block

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const capability = 4

func checkBlock(t *testing.T, b *Block[int], expect []int) {
	require.Equal(t, len(expect), b.Size())
	require.Equal(t, len(expect) == 0, b.IsEmpty())
	require.Equal(t, len(expect) == capability, b.IsFull())
	if len(expect) > 0 {
		require.Equal(t, expect[0], b.Front())
		require.Equal(t, expect[len(expect)-1], b.Back())
	}
	i := 0
	for e := range b.Forward() {
		require.Equal(t, expect[i], e)
		i++
	}
	i = len(expect) - 1
	for e := range b.Backward() {
		require.Equal(t, expect[i], e)
		i--
	}
}

func TestBlock(t *testing.T) {
	b := NewBlock[int](capability)
	checkBlock(t, b, []int{})

	require.Equal(t, 1, b.PushFront(1))
	checkBlock(t, b, []int{1})

	require.Equal(t, 2, b.PushBack(2, 3))
	checkBlock(t, b, []int{1, 2, 3})

	require.Equal(t, 1, b.PushFront(4, 5))
	checkBlock(t, b, []int{4, 1, 2, 3})

	require.Equal(t, 3, b.PopBack())
	checkBlock(t, b, []int{4, 1, 2})

	require.Equal(t, 4, b.PopFront())
	checkBlock(t, b, []int{1, 2})

	// break iter
	for e := range b.Forward() {
		if e == 2 {
			break
		}
	}
	for e := range b.Backward() {
		if e == 1 {
			break
		}
	}

	b.Clear()
	checkBlock(t, b, []int{})

	require.Panics(t, func() { b.Front() })
	require.Panics(t, func() { b.Back() })
	require.Panics(t, func() { b.PopFront() })
	require.Panics(t, func() { b.PopBack() })

	require.Equal(t, 1, b.PushBack(1))
	checkBlock(t, b, []int{1})

	require.Equal(t, 3, b.PushBack(2, 3, 4, 5))
	checkBlock(t, b, []int{1, 2, 3, 4})
}
