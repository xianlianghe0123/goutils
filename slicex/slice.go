package slicex

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Pop removes and returns the last element from s.
// It panics if s is empty.
func Pop[S ~[]T, T any](s *S) T {
	tail := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tail
}

// Fill sets all elements of s with the given value e.
// Note: Pointers will be shared. Use FillFunc for unique instances.
func Fill[S ~[]T, T any](s S, e T) {
	FillFunc(s, func() T { return e })
}

// FillFunc initializes each element of s with the value returned by f.
func FillFunc[S ~[]T, T any](s S, f func() T) {
	for i := range s {
		s[i] = f()
	}
}

// LastIndex returns the index of the last occurrence of e in s,
// or -1 if not present.
func LastIndex[S ~[]T, T comparable](s S, e T) int {
	return LastIndexFunc(s, func(ee T) bool { return ee == e })
}

// LastIndexFunc returns the index of the last element that satisfies,
// or -1 if no elements satisfy the predicate.
func LastIndexFunc[S ~[]T, T any](s S, predicate func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if predicate(s[i]) {
			return i
		}
	}
	return -1
}

// Every checks if all elements satisfy the predicate.
// Returns true for empty slices.
func Every[S ~[]T, T any](s S, predicate func(T) bool) bool {
	for i := range s {
		if !predicate(s[i]) {
			return false
		}
	}
	return true
}

// Count returns the number of occurrences of element e.
func Count[S ~[]T, T comparable](s S, e T) int {
	return CountFunc(s, func(ee T) bool { return ee == e })
}

// CountFunc counts elements satisfying the predicate.
func CountFunc[S ~[]T, T any](s S, predicate func(T) bool) int {
	cnt := 0
	for i := range s {
		t := predicate(s[i])
		cnt += int(*(*byte)(unsafe.Pointer(&t)))
	}
	return cnt
}

// Filter returns a new slice containing elements that satisfy predicate.
func Filter[S ~[]T, T any](s S, predicate func(T) bool) S {
	if len(s) == 0 {
		return nil
	}
	ret := make(S, 0)
	for i := range s {
		if predicate(s[i]) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

// FilterInPlace filters elements in-place using predicate f.
// Modifies original slice but preserves underlying array.
func FilterInPlace[S ~[]T, T any](s S, predicate func(T) bool) S {
	ret := s[:0]
	for i := range s {
		if predicate(s[i]) {
			ret = append(ret, s[i])
		}
	}
	return ret
}

// CastNumber converts slice elements between numeric types.
func CastNumber[T2, T constraints.Integer | constraints.Float, S ~[]T](s S) []T2 {
	if len(s) == 0 {
		return nil
	}
	ret := make([]T2, len(s))
	for i, e := range s {
		ret[i] = T2(e)
	}
	return ret
}

// CastString converts between string-like types (string/[]byte).
func CastString[T2, T ~string | ~[]byte, S ~[]T](s S) []T2 {
	if len(s) == 0 {
		return nil
	}
	ret := make([]T2, len(s))
	for i, e := range s {
		ret[i] = T2(e)
	}
	return ret
}

// Map converts elements using function convert, returns a new slice.
func Map[T2, T any, S ~[]T](s S, convert func(T) T2) []T2 {
	return MapFilter(s, func(e T) (T2, bool) { return convert(e), true })
}

// MapFilter converts and filters elements using a convert function, returns a new slice.
func MapFilter[T2, T any, S ~[]T](s S, convert func(T) (T2, bool)) []T2 {
	if len(s) == 0 {
		return nil
	}
	ret := make([]T2, 0, len(s))
	for i := range s {
		if e2, ok := convert(s[i]); ok {
			ret = append(ret, e2)
		}
	}
	return ret
}

// Reduce iteratively combines elements of s into a single value through an accumulation function.
func Reduce[R, T any, S ~[]T](s S, r R, accumulation func(R, T) R) R {
	for i := range s {
		r = accumulation(r, s[i])
	}
	return r
}

// ToMap converts slice to map using f to generate key-value pairs.
func ToMap[K comparable, V, T any, S ~[]T](s S, f func(T) (K, V)) map[K]V {
	if len(s) == 0 {
		return nil
	}
	m := make(map[K]V, len(s))
	for i := range s {
		k, v := f(s[i])
		m[k] = v
	}
	return m
}
