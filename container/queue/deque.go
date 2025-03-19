package queue

import (
	"iter"

	"github.com/xianlianghe0123/goutils/container/block"
)

// Define the default block.Block capacity
var _BlockCapability = 64

type _DequeNode[T any] struct {
	*block.Block[T]
	prev, next *_DequeNode[T]
}

// Deque double-ended queue.
// use NewDeque to create.
type Deque[T any] struct {
	head, tail *_DequeNode[T]
	size       int
}

// NewDeque returns an initialized empty Deque.
func NewDeque[T any]() *Deque[T] {
	d := &Deque[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
	d.head = d.newNode(nil, nil)
	d.tail = d.head
	return d
}

// Size returns the number of elements. O(1)
func (q *Deque[T]) Size() int { return q.size }

// IsEmpty returns whether has elements.
func (q *Deque[T]) IsEmpty() bool { return q.size == 0 }

// PushFront inserts elements to the front.
func (q *Deque[T]) PushFront(es ...T) {
	q.size += len(es)
	for len(es) > 0 {
		// extend when block is full
		if q.head.IsFull() {
			q.head.prev = q.newNode(nil, q.head)
			q.head = q.head.prev
		}
		batchSize := min(_BlockCapability-q.head.Size(), len(es))
		q.head.PushFront(es[:batchSize]...)
		es = es[batchSize:]
	}
}

// PopFront returns the first element, and remove it from deque.
// It panics when Deque is empty.
func (q *Deque[T]) PopFront() T {
	q.Front()
	q.size--
	return q.head.PopFront()
}

// Front returns the first element.
// It panics when Deque is empty.
func (q *Deque[T]) Front() T {
	if q.IsEmpty() {
		panic("deque is empty")
	}
	if q.head.IsEmpty() {
		q.head = q.head.next
		// for garbage collection, unlink the previous node
		q.head.prev.next = nil
		q.head.prev = nil
	}
	return q.head.Front()
}

// PushBack inserts elements to the back.
func (q *Deque[T]) PushBack(es ...T) {
	q.size += len(es)
	for len(es) > 0 {
		// extend when tail block is full
		if q.tail.IsFull() {
			q.tail.next = q.newNode(q.tail, nil)
			q.tail = q.tail.next
		}
		batchSize := min(_BlockCapability-q.tail.Size(), len(es))
		q.tail.PushBack(es[:batchSize]...)
		es = es[batchSize:]
	}
}

// PopBack return the last element，and remove it from deque.
// It panics when Deque is empty.
func (q *Deque[T]) PopBack() T {
	q.Back()
	q.size--
	return q.tail.PopBack()
}

// Back return the last element.
// It panics when Deque is empty.
func (q *Deque[T]) Back() T {
	if q.IsEmpty() {
		panic("it's empty")
	}
	if q.tail.IsEmpty() {
		q.tail = q.tail.prev
		// for gc
		q.tail.next.prev = nil
		q.tail.next = nil
	}
	return q.tail.Back()
}

// Clear empties the deque.
func (q *Deque[T]) Clear() {
	q.tail = q.head
	q.head.Clear()
	q.size = 0
}

// Forward returns an iterator that yields elements from first to last
func (q *Deque[T]) Forward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if q.IsEmpty() {
			return
		}
		for i := q.head; i != nil; i = i.next {
			for v := range i.Forward() {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Backward returns an iterator that yields elements from last to first
func (q *Deque[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if q.IsEmpty() {
			return
		}
		for i := q.tail; i != nil; i = i.prev {
			for v := range i.Backward() {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (_ *Deque[T]) newNode(prev, next *_DequeNode[T]) *_DequeNode[T] {
	return &_DequeNode[T]{
		Block: block.NewBlock[T](_BlockCapability),
		prev:  prev,
		next:  next,
	}
}
