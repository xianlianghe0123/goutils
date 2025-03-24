package tree

type UnionFind[T comparable] struct {
	parent map[T]T
	// connected component count
	cnt int
}

// NewUnionFind creates and returns a new instance of UnionFind.
func NewUnionFind[T comparable]() *UnionFind[T] {
	uf := &UnionFind[T]{
		parent: make(map[T]T),
		cnt:    0,
	}
	return uf
}

// Find returns the root of the element e.
func (uf *UnionFind[T]) Find(e T) T {
	if _, ok := uf.parent[e]; !ok {
		uf.parent[e] = e
		uf.cnt++
	}
	if uf.parent[e] != e {
		uf.parent[e] = uf.Find(uf.parent[e])
	}
	return uf.parent[e]
}

// IsConnect checks if two elements are in the same set.
func (uf *UnionFind[T]) IsConnect(e1, e2 T) bool { return uf.Find(e1) == uf.Find(e2) }

// Union merges the sets containing elements e1 and e2.
func (uf *UnionFind[T]) Union(e1, e2 T) {
	if uf.IsConnect(e1, e2) {
		return
	}
	uf.parent[uf.Find(e1)] = uf.Find(e2)
	uf.cnt--
}

// ConnectedComponent returns the number of connected components in the union-find.
func (uf *UnionFind[T]) ConnectedComponent() int { return uf.cnt }
