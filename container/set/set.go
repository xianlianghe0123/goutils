package set

import (
	"iter"
	"maps"
)

type Set[T comparable] map[T]struct{}

// NewSet creates a new empty Set of type T.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Size returns the number of elements.
func (s Set[T]) Size() int { return len(s) }

// Add adds elements to the Set.
func (s Set[T]) Add(es ...T) {
	for i := range es {
		s[es[i]] = struct{}{}
	}
}

// Remove removes multiple elements from the set.
func (s Set[T]) Remove(es ...T) {
	for i := range es {
		delete(s, es[i])
	}
}

// Contains checks if an element exists in the Set.
func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

// Clear removes all elements from the Set.
func (s Set[T]) Clear() { clear(s) }

// Union returns a new Set that is the union of the current Set and another Set.
// The union of two sets contains all elements that are in either Set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	ret := maps.Clone(s)
	maps.Copy(ret, other)
	return ret
}

// Intersection returns a new Set that is the intersection of the current Set and another Set.
// The intersection of two sets contains all elements that are in both sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	less, more := s, other
	if less.Size() > more.Size() {
		less, more = more, less
	}
	ret := make(Set[T], less.Size())
	for e := range less {
		if more.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Difference returns a new Set that is the difference between the current Set and another Set.
// The difference of two sets contains all elements that are in the current Set but not in the other Set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	ret := make(Set[T], s.Size()-other.Size())
	for e := range s {
		if !other.Contains(e) {
			ret.Add(e)
		}
	}
	return ret
}

// Iter returns an iterator to iterate over the elements of the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range s {
			if !yield(e) {
				return
			}
		}
	}
}
