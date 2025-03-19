package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkDeque(t *testing.T, d *Deque[int], expect []int) {
	require.Equal(t, len(expect), d.Size())
	require.Equal(t, len(expect) == 0, d.IsEmpty())
	if len(expect) > 0 {
		require.Equal(t, expect[0], d.Front())
		require.Equal(t, expect[len(expect)-1], d.Back())
	}
	i := 0
	for e := range d.Forward() {
		require.Equal(t, expect[i], e)
		i++
	}
	i = len(expect) - 1
	for e := range d.Backward() {
		require.Equal(t, expect[i], e)
		i--
	}
}

func TestDeque(t *testing.T) {
	_BlockCapability = 4
	d := NewDeque[int]()

	d.PushBack(1, 2, 3, 4, 5)
	checkDeque(t, d, []int{1, 2, 3, 4, 5})

	d.PushFront(6, 7, 8, 9, 10)
	checkDeque(t, d, []int{10, 9, 8, 7, 6, 1, 2, 3, 4, 5})

	require.Equal(t, 10, d.PopFront(), 10)
	require.Equal(t, 9, d.PopFront(), 9)
	checkDeque(t, d, []int{8, 7, 6, 1, 2, 3, 4, 5})

	require.Equal(t, 5, d.PopBack())
	require.Equal(t, 4, d.PopBack())
	require.Equal(t, 3, d.PopBack())
	require.Equal(t, 2, d.PopBack())
	checkDeque(t, d, []int{8, 7, 6, 1})

	// check iter break
	for e := range d.Forward() {
		if e == 7 {
			break
		}
	}
	for e := range d.Backward() {
		if e == 6 {
			break
		}
	}

	d.Clear()
	checkDeque(t, d, []int{})

	require.Panics(t, func() { d.Front() })
	require.Panics(t, func() { d.Back() })
	require.Panics(t, func() { d.PopFront() })
	require.Panics(t, func() { d.PopBack() })
}

func BenchmarkDeque(b *testing.B) {
	_BlockCapability = 64
	for range b.N {
		d := NewDeque[int]()
		for i := range int(1e4) {
			d.PushBack(i)
		}
		for d.Size() > 0 {
			d.PopFront()
		}
	}
}
