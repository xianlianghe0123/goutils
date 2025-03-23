package stream

import (
	"iter"
)

// Map transforms a stream of type T into a stream of type R.
func Map[R, T any](s Stream[T], f func(T) R) Stream[R] {
	return func(yield func(R) bool) {
		for e := range s {
			if !yield(f(e)) {
				return
			}
		}
	}
}

// FlatMap transforms a stream of type T into a stream of type R.
// The new stream consisting of the results of replacing each element of
// this stream with the contents of a mapped stream produced by applying
func FlatMap[R, T any](s Stream[T], f func(T) iter.Seq[R]) Stream[R] {
	return func(yield func(R) bool) {
		for e := range s {
			for v := range f(e) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
