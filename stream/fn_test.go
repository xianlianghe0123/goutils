package stream

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	s := NewStream(slices.Values([]int{2, 1, 3}))
	sr := Map(s, func(e int) string {
		return strconv.Itoa(e)
	})
	require.Equal(t, []string{"2", "1", "3"}, sr.Collect())
	for range sr {
		break
	}
}

func TestFlatMap(t *testing.T) {
	s := NewStream(slices.Values([]int{2, 1, 3}))
	sr := FlatMap(s, func(e int) iter.Seq[string] {
		return func(yield func(string) bool) {
			for i := 0; i < e; i++ {
				if !yield(strconv.Itoa(e)) {
					return
				}
			}
		}
	})
	require.Equal(t, []string{"2", "2", "1", "3", "3", "3"}, sr.Collect())
	for range sr {
		break
	}
}
