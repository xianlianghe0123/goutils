package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie(t *testing.T) {
	trie := NewTrie[byte]()
	trie.Add(3, []byte("apple")...)
	trie.Add(2, []byte("banana")...)
	trie.Add(1, []byte("cherry")...)
	trie.Add(1, []byte("peach")...)
	trie.Add(1, []byte("pear")...)
	trie.Add(1, []byte("pineapple")...)
	// remove prefix
	trie.Remove(1, []byte("app")...)
	// remove partially
	trie.Remove(1, []byte("apple")...)
	// remove all
	trie.Remove(1, []byte("peach")...)
	trie.Remove(0, []byte("banana")...)
	// remove not exist
	trie.Remove(0, []byte("orange")...)

	// check partially
	require.True(t, trie.HasPrefix([]byte("app")...))
	// check removed
	require.False(t, trie.HasPrefix([]byte("banana")...))
	require.False(t, trie.HasPrefix([]byte("peac")...))
	// check cherry
	require.True(t, trie.HasPrefix([]byte("cherry")...))
	// check not exist
	require.False(t, trie.HasPrefix([]byte("orange")...))

	// exist
	require.Equal(t, 2, trie.Find([]byte("apple")...))
	require.Equal(t, 1, trie.Find([]byte("cherry")...))
	// check removed
	require.Equal(t, 0, trie.Find([]byte("banana")...))
	require.Equal(t, 0, trie.Find([]byte("peach")...))
	// not exist
	require.Equal(t, 0, trie.Find([]byte("orange")...))

	// has data
	m := make(map[string]int)
	trie.ForEach(func(data []byte, cnt int) {
		m[string(data)] = cnt
	})
	require.Equal(t, map[string]int{
		"apple":     2,
		"cherry":    1,
		"pear":      1,
		"pineapple": 1,
	}, m)
	// prefix exist
	m = make(map[string]int)
	trie.ForEachPrefix([]byte("p"), func(data []byte, cnt int) {
		m[string(data)] = cnt
	})
	require.Equal(t, map[string]int{
		"pear":      1,
		"pineapple": 1,
	}, m)
	// prefix not exist
	cnt := 0
	trie.ForEachPrefix([]byte("x"), func(data []byte, cnt int) { cnt++ })
	require.Equal(t, 0, cnt)
}
