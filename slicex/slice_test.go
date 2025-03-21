package slicex

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	for i := 5; i > 0; i-- {
		require.Equal(t, i, Pop(&s))
	}
	require.Panics(t, func() { Pop(&s) })
}

func TestFill(t *testing.T) {
	s := []int{0, 0, 0}
	Fill(s, 10)
	require.Equal(t, []int{10, 10, 10}, s)
}

func TestFillFunc(t *testing.T) {
	s := []int{0, 0, 0}
	FillFunc(s, func() int { return 10 })
	require.Equal(t, []int{10, 10, 10}, s)
}

func TestLastIndex(t *testing.T) {
	s := []int{1, 1, 2, 2, 1, 3, 4}
	require.Equal(t, 4, LastIndex(s, 1))
	require.Equal(t, -1, LastIndex(s, 100))
}

func TestLastIndexFunc(t *testing.T) {
	s := []int{1, 1, 2, 2, 1, 3, 4}
	require.Equal(t, 4, LastIndexFunc(s, func(e int) bool { return e == 1 }))
	require.Equal(t, -1, LastIndexFunc(s, func(e int) bool { return e == 100 }))
}

func TestEvery(t *testing.T) {
	s := []int{1, 1, 2, 2, 1, 3, 4}
	require.True(t, Every(s, func(e int) bool { return e > 0 }))
	require.False(t, Every(s, func(e int) bool { return e < 3 }))
}

func TestCount(t *testing.T) {
	require.Equal(t, 3, Count([]int{1, 2, 2, 3, 3, 3, 4}, 3))
}

func TestCountFunc(t *testing.T) {
	require.Equal(t, 3,
		CountFunc([]int{1, 2, 3, 4, 5}, func(e int) bool { return e%2 == 1 }))
}

func TestFilter(t *testing.T) {
	require.Empty(t, Filter([]int{}, func(e int) bool { return true }))
	require.Equal(t, []int{2, 4, 6},
		Filter([]int{1, 2, 3, 4, 5, 6, 7}, func(e int) bool { return e%2 == 0 }))
}

func TestFilterInPlace(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	ss := FilterInPlace(s, func(e int) bool {
		return e%2 == 0
	})
	require.Equal(t, s[:len(ss)], ss)
}

func TestCastNumber(t *testing.T) {
	require.Empty(t, CastNumber[int]([]int16{}), 0)
	require.Equal(t, []int{1, 2, 3}, CastNumber[int]([]int16{1, 2, 3}))
}

func TestCastString(t *testing.T) {
	require.Empty(t, CastString[[]byte]([]string{}))
	require.Equal(t, []string{"a", "b", "c"},
		CastString[string]([][]byte{{'a'}, {'b'}, {'c'}}))
}

func TestMap(t *testing.T) {
	require.Empty(t, Map([]int{}, func(e int) int { return e }))
	require.Equal(t, []string{"1", "2", "3"},
		Map([]int{1, 2, 3}, func(e int) string { return strconv.Itoa(e) }))
}

func TestMapFilter(t *testing.T) {
	require.Empty(t, MapFilter([]int{}, func(e int) (int, bool) { return e, true }))
	require.Equal(t, []string{"1", "3"},
		MapFilter([]int{1, 2, 3}, func(e int) (string, bool) {
			return strconv.Itoa(e), e%2 == 1
		}))
}

func TestReduce(t *testing.T) {
	require.Equal(t, 15,
		Reduce([]int{1, 2, 3, 4, 5}, 0, func(a, b int) int { return a + b }))
}

func TestToMap(t *testing.T) {
	require.Empty(t, ToMap([]int{}, func(e int) (int, int) { return e, e }))
	require.Equal(t, map[string]int{"1": 2, "2": 3},
		ToMap([]int{1, 2}, func(e int) (string, int) { return strconv.Itoa(e), e + 1 }))
}

func BenchmarkFill(b *testing.B) {
	s := make([]int, 1e4)
	for i := range b.N {
		Fill(s, i)
	}
}

func BenchmarkFillFunc(b *testing.B) {
	s := make([]int, 1e4)
	for i := range b.N {
		FillFunc(s, func() int {
			return i
		})
	}
}
