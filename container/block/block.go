package block

import (
	"iter"
)

// Block is a circular slice with fixed size.
// Use NewBlock to create.
type Block[T any] struct {
	elems      []T // store the elements
	head, tail int // record the first and last index of elements
	size       int // the number of elements
}

// NewBlock returns an initialized empty block.
func NewBlock[T any](capacity int) *Block[T] {
	return &Block[T]{
		elems: make([]T, capacity),
		head:  0,
		tail:  capacity - 1,
	}
}

// Size returns the number of elements. O(1)
func (r *Block[T]) Size() int { return r.size }

// IsEmpty returns whether the block has no elements.
func (r *Block[T]) IsEmpty() bool { return r.size == 0 }

// IsFull returns whether the block is full
func (r *Block[T]) IsFull() bool { return r.size == cap(r.elems) }

// PushFront inserts elements to the front and returns the number of insertion.
// When the number of elements greater than capability, ignore the redundant elements.
func (r *Block[T]) PushFront(elems ...T) int {
	cnt := min(cap(r.elems)-r.size, len(elems))
	for i := range cnt {
		r.head = r.decr(r.head)
		r.elems[r.head] = elems[i]
	}
	r.size += cnt
	return cnt
}

// PopFront removes and returns the first element.
// It panics when block is empty.
func (r *Block[T]) PopFront() T {
	if r.IsEmpty() {
		panic("block is empty")
	}
	ret := r.elems[r.head]
	r.head = r.incr(r.head)
	r.size--
	return ret
}

// Front return the first element.
// It panics when block is empty.
func (r *Block[T]) Front() T {
	if r.IsEmpty() {
		panic("block is empty")
	}
	return r.elems[r.head]
}

// PushBack inserts elements to the back and returns the number of insertion.
// When the number of elements greater than capability, ignore the redundant elements.
func (r *Block[T]) PushBack(es ...T) int {
	cnt := min(cap(r.elems)-r.size, len(es))
	for i := range cnt {
		r.tail = r.incr(r.tail)
		r.elems[r.tail] = es[i]
	}
	r.size += cnt
	return cnt
}

// PopBack removes and returns the last element.
// It panics when block is empty.
func (r *Block[T]) PopBack() T {
	if r.IsEmpty() {
		panic("it's empty")
	}
	ret := r.elems[r.tail]
	r.tail = r.decr(r.tail)
	r.size--
	return ret
}

// Back return the last element.
// It panics when block is empty.
func (r *Block[T]) Back() T {
	if r.IsEmpty() {
		panic("it's empty")
	}
	return r.elems[r.tail]
}

// Forward returns an iterator which elements from first to last.
func (r *Block[T]) Forward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if r.IsEmpty() {
			return
		}
		for i := r.head; ; i = r.incr(i) {
			if !yield(r.elems[i]) {
				return
			}
			if i == r.tail {
				break
			}
		}
	}
}

// Backward returns an iterator which elements from last to first.
func (r *Block[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if r.IsEmpty() {
			return
		}
		for i := r.tail; ; i = r.decr(i) {
			if !yield(r.elems[i]) {
				return
			}
			if i == r.head {
				break
			}
		}
	}
}

// Clear empties the Block.
func (r *Block[T]) Clear() {
	r.head, r.tail = 0, cap(r.elems)-1
	r.size = 0
}

func (r *Block[T]) incr(i int) int { return (i + 1) % cap(r.elems) }

func (r *Block[T]) decr(i int) int { return (i - 1 + cap(r.elems)) % cap(r.elems) }
