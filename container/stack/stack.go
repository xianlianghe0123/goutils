package stack

import (
	"iter"

	"github.com/xianlianghe0123/goutils/slicex"
)

type Stack[T any] struct {
	elems []T
}

// NewStack create a Stack with an initialize capacity
func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0, capacity),
	}
}

// Size returns the number of elements.
func (s *Stack[T]) Size() int { return len(s.elems) }

// IsEmpty checks if the Stack is empty.
func (s *Stack[T]) IsEmpty() bool { return s.Size() == 0 }

// Push adds an element to the top of the Stack.
func (s *Stack[T]) Push(es ...T) { s.elems = append(s.elems, es...) }

// Pop removes and returns the top element from the Stack.
// It panics if the Stack is empty.
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return slicex.Pop(&s.elems)
}

// Top returns the top element of the Stack.
// It panics if the Stack is empty.
func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		panic("it's empty")
	}
	return s.elems[s.Size()-1]
}

// Clear empties the Stack.
func (s *Stack[T]) Clear() { s.elems = s.elems[:0] }

// Iter returns an iterator that traverses the Stack from top to bottom.
func (s *Stack[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := s.Size() - 1; i >= 0; i-- {
			if !yield(s.elems[i]) {
				return
			}
		}
	}
}
