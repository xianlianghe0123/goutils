package stream

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStream_Filter(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3, 4, 5})).
		Filter(func(n int) bool {
			return n%2 == 0
		})
	require.Equal(t, []int{2, 4}, s.Collect())
	for range s {
		break
	}
}

func TestStream_Skip(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3, 4, 5})).
		Skip(3)
	require.Equal(t, []int{4, 5}, s.Collect())
	require.Equal(t, []int{4, 5}, s.Collect())
	for range s {
		break
	}
}

func TestStream_Limit(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3, 4, 5})).
		Limit(3)
	require.Equal(t, []int{1, 2, 3}, s.Collect())
	require.Equal(t, []int{1, 2, 3}, s.Collect())
	for range s {
		break
	}
}

func TestStream_Distinct(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 1, 2, 2, 3})).
		Distinct(func(n int) string {
			return strconv.Itoa(n)
		})
	require.Equal(t, []int{1, 2, 3}, s.Collect())
	for range s {
		break
	}
}

func TestStream_Sort(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3, 4, 5})).
		Sort(func(i int, i2 int) int {
			return i2 - i
		})
	require.Equal(t, []int{5, 4, 3, 2, 1}, s.Collect())
	for range s {
		break
	}
}

func TestStream_ForEach(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3})).
		ForEach(func(i int) int {
			return i + 1
		})
	require.Equal(t, []int{2, 3, 4}, s.Collect())
	for range s {
		break
	}
}

func TestStream_Iter(t *testing.T) {
	iter := NewStream(slices.Values([]int{1, 2, 3})).
		Iter()
	require.Equal(t, []int{1, 2, 3}, slices.Collect(iter))
}

func TestStream_Collect(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3})).
		Collect()
	require.Equal(t, []int{1, 2, 3}, s)
}

func TestStream_Count(t *testing.T) {
	c := NewStream(slices.Values([]int{1, 2, 3})).
		Count()
	require.Equal(t, 3, c)
}

func TestStream_Every(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3}))
	require.True(t, s.Every(func(i int) bool { return i > 0 }))
	require.False(t, s.Every(func(i int) bool { return i == 2 }))
}

func TestStream_Some(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3}))
	require.True(t, s.Some(func(i int) bool { return i == 2 }))
	require.False(t, s.Some(func(i int) bool { return i < 0 }))
}

func TestStream_Find(t *testing.T) {
	s := NewStream(slices.Values([]int{2, 1, 3}))
	ret, ok := s.Find(func(i int) bool { return i == 2 })
	require.True(t, ok)
	require.Equal(t, 2, ret)

	ret, ok = s.Find(func(i int) bool { return i == 0 })
	require.False(t, ok)
	require.Equal(t, 0, ret)
}

func TestStream_Min(t *testing.T) {
	s := NewStream(slices.Values([]int{2, 1, 3}))
	ret, ok := s.Min(func(i, i2 int) int { return i - i2 })
	require.True(t, ok)
	require.Equal(t, 1, ret)

	s = NewStream(slices.Values([]int{}))
	ret, ok = s.Min(func(i, i2 int) int { return i - i2 })
	require.False(t, ok)
	require.Equal(t, 0, ret)
}

func TestStream_Max(t *testing.T) {
	s := NewStream(slices.Values([]int{1, 2, 3}))
	ret, ok := s.Max(func(i, i2 int) int { return i - i2 })
	require.True(t, ok)
	require.Equal(t, 3, ret)

	s = NewStream(slices.Values([]int{}))
	ret, ok = s.Max(func(i, i2 int) int { return i - i2 })
	require.False(t, ok)
	require.Equal(t, 0, ret)
}
