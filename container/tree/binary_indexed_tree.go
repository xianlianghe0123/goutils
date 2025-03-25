package tree

// BinaryIndexedTree represents a Binary Indexed Tree (Fenwick Tree) data structure.
type BinaryIndexedTree[T any] struct {
	origin []T
	tree   []T
	merge  func(T, T) T
}

// NewBinaryIndexedTree creates and initials a new Binary Indexed Tree
// from a given slice of elements.
// @Param merge defines how to merge two elements.
// @Complexity O(n)
func NewBinaryIndexedTree[T any](origin []T, merge func(x T, y T) T) *BinaryIndexedTree[T] {
	bit := &BinaryIndexedTree[T]{
		origin: origin,
		tree:   append([]T(nil), origin...),
		merge:  merge,
	}
	for i := 1; i <= len(origin); i++ {
		if j := i + lowBit(i); j <= len(bit.tree) {
			bit.tree[j-1] = merge(bit.tree[j-1], bit.tree[i-1])
		}
	}
	return bit
}

// Accumulate calculates the prefix accumulated value up to the given index i.
// @Complexity O(log(n))
func (bit *BinaryIndexedTree[T]) Accumulate(i int) T {
	ret := bit.tree[i]
	for i = (i + 1) & i; i > 0; i &= i - 1 {
		ret = bit.merge(ret, bit.tree[i-1])
	}
	return ret
}

// Renew updates the element at the given index i.
// @Param change returns the new element from old.
// @Complexity O(log(n))
func (bit *BinaryIndexedTree[T]) Renew(i int, change func(T) T) {
	bit.origin[i] = change(bit.origin[i])
	for i++; i <= len(bit.tree); i += lowBit(i) {
		bit.tree[i-1] = change(bit.tree[i-1])
	}
}

// Range consumes the accumulated value within the range [l, r].
// @Complexity O(log^2(n))
func (bit *BinaryIndexedTree[T]) Range(l, r int, consume func(e T)) {
	l, r = l+1, r+1
	for r >= l {
		if r-lowBit(r)+1 >= l {
			consume(bit.tree[r-1])
			r &= r - 1
		} else {
			consume(bit.origin[r-1])
			r--
		}
	}
}

// GetOrigin returns the original element at the given index i.
func (bit *BinaryIndexedTree[T]) GetOrigin(i int) T { return bit.origin[i] }

// lowBit calculates the lowest set bit of a given integer i.
func lowBit(i int) int { return i & -i }
