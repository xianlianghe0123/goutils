package queue

import (
	"github.com/xianlianghe0123/goutils/slicex"
)

type PriorityQueue[T any] struct {
	elems   []T
	_higher func(T, T) bool
}

// NewPriorityQueue returns an empty PriorityQueue.
// @Param higher is a function that compares two elements of type T.
// It should return true if t1 has a higher priority than t2.
func NewPriorityQueue[T any](higher func(T, T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		elems:   make([]T, 0),
		_higher: higher,
	}
}

// Size returns the number of elements.
func (pq *PriorityQueue[T]) Size() int { return len(pq.elems) }

// IsEmpty returns true when has no elements, otherwise false.
func (pq *PriorityQueue[T]) IsEmpty() bool { return pq.Size() == 0 }

// Init use the elements initialize the PriorityQueue.
func (pq *PriorityQueue[T]) Init(elems ...T) {
	pq.elems = append(pq.elems[:0], elems...)
	for i := pq.Size()/2 - 1; i >= 0; i-- {
		pq.down(i)
	}
}

// Push add an element to the PriorityQueue.
func (pq *PriorityQueue[T]) Push(e T) {
	pq.elems = append(pq.elems, e)
	pq.up(pq.Size() - 1)
}

// Pop returns the highest priority element and remove it.
// It panics when the PriorityQueue is empty.
func (pq *PriorityQueue[T]) Pop() T {
	if pq.IsEmpty() {
		panic("PriorityQueue is empty")
	}
	pq.swap(0, pq.Size()-1)
	e := slicex.Pop(&pq.elems)
	pq.down(0)
	return e
}

// Front returns the highest priority element.
// It panics when the PriorityQueue is empty.
func (pq *PriorityQueue[T]) Front() T {
	if pq.IsEmpty() {
		panic("PriorityQueue is empty")
	}
	return pq.elems[0]
}

// Clear empties the PriorityQueue.
func (pq *PriorityQueue[T]) Clear() { pq.elems = pq.elems[:0] }

func (pq *PriorityQueue[T]) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		// break early if the elements have the same priority
		if !pq.higher(i, parent) {
			break
		}
		pq.swap(i, parent)
		i = parent
	}
}

func (pq *PriorityQueue[T]) down(i int) {
	for {
		left, right := 2*i+1, 2*i+2
		if left >= pq.Size() {
			break
		}
		t := left
		if right < pq.Size() && pq.higher(right, left) {
			t = right
		}
		// break earlier when has same priority
		if !pq.higher(t, i) {
			break
		}
		pq.swap(i, t)
		i = t
	}
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.elems[i], pq.elems[j] = pq.elems[j], pq.elems[i]
}

func (pq *PriorityQueue[T]) higher(i, j int) bool {
	return pq._higher(pq.elems[i], pq.elems[j])
}
