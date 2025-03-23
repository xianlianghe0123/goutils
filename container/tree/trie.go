package tree

type _TrieNode[T comparable] struct {
	cnt      int
	children map[T]*_TrieNode[T]
}

// Trie is a generic prefix tree data structure.
type Trie[T comparable] struct {
	root *_TrieNode[T]
}

// NewTrie creates a new instance of the Trie tree.
func NewTrie[T comparable]() *Trie[T] {
	return &Trie[T]{
		root: new(_TrieNode[T]),
	}
}

// Add inserts a sequence of elements into the Trie with a given count.
func (t *Trie[T]) Add(cnt int, seq ...T) {
	n := t.root
	for i := range seq {
		if n.children[seq[i]] == nil {
			if n.children == nil {
				n.children = make(map[T]*_TrieNode[T])
			}
			n.children[seq[i]] = new(_TrieNode[T])
		}
		n = n.children[seq[i]]
	}
	n.cnt += cnt
}

// Remove deletes a sequence of elements from the Trie with a given count.
// If the count is less than or 0, the entire sequence is removed.
// Or the count is decremented. After that, if the sequence count is less than or 0,
// the entire sequence is removed too.
func (t *Trie[T]) Remove(cnt int, seq ...T) {
	nodes := make([]*_TrieNode[T], 0, len(seq))
	n := t.root
	for i := range seq {
		if n.children[seq[i]] == nil {
			return
		}
		nodes = append(nodes, n)
		n = n.children[seq[i]]
	}
	if n == nil || n.cnt == 0 {
		return
	}
	// remove partially
	if 0 < cnt && cnt < n.cnt {
		n.cnt -= cnt
		return
	}
	// remove all
	n.cnt = 0
	for i := len(seq) - 1; i >= 0; i-- {
		if len(n.children) > 0 || n.cnt > 0 {
			return
		}
		delete(nodes[i].children, seq[i])
		n = nodes[i]
	}
}

// Find searches for a sequence of elements in the Trie and returns its count.
func (t *Trie[T]) Find(seq ...T) int {
	n := t.find(seq)
	if n == nil {
		return 0
	}
	return n.cnt
}

// HasPrefix checks if the Trie has the prefix.
func (t *Trie[T]) HasPrefix(prefix ...T) bool {
	n := t.find(prefix)
	return n != nil
}

// ForEach iterates over all sequences in the Trie and applies a consumer function.
func (t *Trie[T]) ForEach(consume func(seq []T, cnt int)) {
	t.forEach(t.root, nil, consume)
}

// ForEachPrefix iterates over all sequences in the Trie that start with a given prefix.
func (t *Trie[T]) ForEachPrefix(prefix []T, consume func(seq []T, cnt int)) {
	n := t.find(prefix)
	t.forEach(n, prefix, consume)
}

// find returns the last node satisfy the prefix.
func (t *Trie[T]) find(prefix []T) *_TrieNode[T] {
	n := t.root
	for i := range prefix {
		if n.children[prefix[i]] == nil {
			return nil
		}
		n = n.children[prefix[i]]
	}
	return n
}

// forEach recursively traverses the Trie tree and applies a consumer function.
func (t *Trie[T]) forEach(n *_TrieNode[T], prefix []T, consume func(seq []T, cnt int)) {
	if n == nil {
		return
	}
	if n.cnt > 0 {
		consume(prefix, n.cnt)
	}
	for e, child := range n.children {
		t.forEach(child, append(prefix, e), consume)
	}
}
