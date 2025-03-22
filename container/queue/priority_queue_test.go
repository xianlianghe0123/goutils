package queue

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/xianlianghe0123/goutils/mathx"
)

func checkPriorityQueue(t *testing.T, q *PriorityQueue[int], size int) {
	require.Equal(t, size, q.Size())
	require.Equal(t, size == 0, q.IsEmpty())
	if q.Size() > 1 {
		q := &PriorityQueue[int]{
			elems:   slices.Clone(q.elems),
			_higher: q._higher,
		}
		e := q.Pop()
		for q.Size() > 0 {
			require.False(t, q._higher(q.Front(), e))
			e = q.Pop()
		}
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int](cmp.Less)
	pq.Init(9, 7, 8, 6, 5, 4, 1, 2, 3)
	checkPriorityQueue(t, pq, 9)

	pq.Push(99)
	checkPriorityQueue(t, pq, 10)

	require.Equal(t, 1, pq.Pop())
	checkPriorityQueue(t, pq, 9)

	pq.Push(0)
	checkPriorityQueue(t, pq, 10)

	require.Equal(t, 0, pq.Pop())
	checkPriorityQueue(t, pq, 9)

	pq.Clear()
	checkPriorityQueue(t, pq, 0)

	require.Panics(t, func() { pq.Front() })
	require.Panics(t, func() { pq.Pop() })
}

func buildSlice(start, end, step int) []int {
	t := make([]int, mathx.Abs(end-start))
	for i := start; i != end; i += step {
		t = append(t, i)
	}
	return t
}

func BenchmarkPriorityQueue_PushAllSame(b *testing.B) {
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		for range int(1e4) {
			pq.Push(0)
		}
	}
}

func BenchmarkPriorityQueue_PushBest(b *testing.B) {
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		for i := range int(1e4) {
			pq.Push(i)
		}
	}
}

func BenchmarkPriorityQueue_PushWorst(b *testing.B) {
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		for i := range int(1e4) {
			pq.Push(1e4 - i)
		}
	}
}

func BenchmarkPriorityQueue_InitAllSame(b *testing.B) {
	t := make([]int, 1e4)
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		pq.Init(t...)
	}
}

func BenchmarkPriorityQueue_InitBest(b *testing.B) {
	t := buildSlice(0, 1e4, 1)
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		pq.Init(t...)
	}
}

func BenchmarkPriorityQueue_InitWorst(b *testing.B) {
	t := buildSlice(1e4, 0, -1)
	for range b.N {
		pq := NewPriorityQueue[int](cmp.Less)
		pq.Init(t...)
	}
}

func BenchmarkPriorityQueue_Pop(b *testing.B) {
	t := buildSlice(1, 1e4, 1)
	for range b.N {
		pq := PriorityQueue[int]{
			elems:   slices.Clone(t),
			_higher: cmp.Less[int],
		}
		for pq.Size() > 0 {
			pq.Pop()
		}
	}
}

func BenchmarkPriorityQueue_PopAllSame(b *testing.B) {
	t := make([]int, 1e4)
	for range b.N {
		pq := PriorityQueue[int]{
			elems:   slices.Clone(t),
			_higher: cmp.Less[int],
		}
		pq.Init(t...)
		for pq.Size() > 0 {
			pq.Pop()
		}
	}
}
