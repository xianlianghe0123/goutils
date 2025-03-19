package queue

import (
	"iter"
)

type Queue[T any] struct {
	deque *Deque[T]
}

// NewQueue creates an empty Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		deque: NewDeque[T](),
	}
}

// Size returns the number of elements.
func (q *Queue[T]) Size() int { return q.deque.Size() }

// IsEmpty returns true when there are no elements, otherwise false.
func (q *Queue[T]) IsEmpty() bool { return q.Size() == 0 }

// Push adds elements to the Queue.
func (q *Queue[T]) Push(es ...T) { q.deque.PushBack(es...) }

// Pop removes and returns the first element.
func (q *Queue[T]) Pop() T { return q.deque.PopFront() }

// Front returns the first element.
// It panics when the Queue is empty.
func (q *Queue[T]) Front() T { return q.deque.Front() }

// Clear empties the Queue.
func (q *Queue[T]) Clear() { q.deque.Clear() }

// Iter returns an iterator that yields elements from first to last.
func (q *Queue[T]) Iter() iter.Seq[T] { return q.deque.Forward() }
