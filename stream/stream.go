package stream

import (
	"iter"
	"slices"

	"github.com/xianlianghe0123/goutils/container/set"
)

// Stream represents a stream of data of type T.
type Stream[T any] iter.Seq[T]

// NewStream creates a new Stream instance from a given iter.Seq[T].
func NewStream[T any](seq iter.Seq[T]) Stream[T] {
	return Stream[T](seq)
}

// Filter filters the elements in the stream based on a given predicate.
func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	return func(yield func(T) bool) {
		for e := range s {
			if !predicate(e) {
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// Skip skips the first n elements in the stream.
func (s Stream[T]) Skip(n int) Stream[T] {
	return func(yield func(T) bool) {
		n := n
		for e := range s {
			if n > 0 {
				n--
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// Limit limits the number of elements in the stream to n.
func (s Stream[T]) Limit(n int) Stream[T] {
	return func(yield func(T) bool) {
		n := n
		for e := range s {
			if !yield(e) {
				return
			}
			n--
			if n == 0 {
				return
			}
		}
	}
}

// Distinct removes duplicate elements from the stream based on a given uniqueness function.
func (s Stream[T]) Distinct(uniq func(T) string) Stream[T] {
	return func(yield func(T) bool) {
		t := set.NewSet[string]()
		for e := range s {
			u := uniq(e)
			if t.Contains(u) {
				continue
			}
			t.Add(u)
			if !yield(e) {
				return
			}
		}
	}
}

// Sort sorts the elements in the stream based on a given comparison function.
func (s Stream[T]) Sort(cmp func(T, T) int) Stream[T] {
	elems := s.Collect()
	slices.SortFunc(elems, cmp)
	return func(yield func(T) bool) {
		for i := range elems {
			if !yield(elems[i]) {
				return
			}
		}
	}
}

// ForEach applies a given function to each element in the stream and
// returns a new stream with the transformed elements.
func (s Stream[T]) ForEach(consume func(T) T) Stream[T] {
	return func(yield func(T) bool) {
		for e := range s {
			if !yield(consume(e)) {
				return
			}
		}
	}
}

// Iter converts the Stream[T] back to an iter.Seq[T].
func (s Stream[T]) Iter() iter.Seq[T] { return iter.Seq[T](s) }

// Collect collects all elements from the stream into a slice.
func (s Stream[T]) Collect() []T { return slices.Collect(iter.Seq[T](s)) }

// Count counts the number of elements in the stream.
func (s Stream[T]) Count() int {
	cnt := 0
	for range s {
		cnt++
	}
	return cnt
}

// Every checks if all elements in the stream satisfy a given predicate
func (s Stream[T]) Every(predicate func(T) bool) bool {
	for e := range s {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// Some checks if at least one element in the stream satisfies a given predicate.
func (s Stream[T]) Some(predicate func(T) bool) bool {
	for e := range s {
		if predicate(e) {
			return true
		}
	}
	return false
}

// Find searches for the first element in the stream that satisfies a given predicate.
func (s Stream[T]) Find(predicate func(T) bool) (ret T, ok bool) {
	for e := range s {
		if predicate(e) {
			return e, true
		}
	}
	return ret, false
}

// Min finds the minimum element in the stream based on a given comparison function.
func (s Stream[T]) Min(cmp func(T, T) int) (ret T, ok bool) {
	for e := range s {
		if !ok {
			ret = e
			ok = true
			continue
		}
		if cmp(e, ret) < 0 {
			ret = e
		}
	}
	return ret, ok
}

// Max finds the maximum element in the stream based on a given comparison function.
func (s Stream[T]) Max(cmp func(T, T) int) (ret T, ok bool) {
	for e := range s {
		if !ok {
			ret = e
			ok = true
			continue
		}
		if cmp(e, ret) > 0 {
			ret = e
		}
	}
	return ret, ok
}
