package tree

import (
	"iter"
)

type _AvlTreeNode[K, V any] struct {
	key K
	val V
	// height means the max height of subtree that this node is the root
	height int
	// size is the subtree's node count
	size        int
	left, right *_AvlTreeNode[K, V]
}

func newNode[K, V any](key K, val V) *_AvlTreeNode[K, V] {
	return &_AvlTreeNode[K, V]{
		key:    key,
		val:    val,
		height: 1,
		size:   1,
	}
}

func (n *_AvlTreeNode[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// getFactor returns the difference between left child height and right child height
func (n *_AvlTreeNode[K, V]) getFactor() int { return n.left.getHeight() - n.right.getHeight() }

func (n *_AvlTreeNode[K, V]) reBalance() *_AvlTreeNode[K, V] {
	switch n.getFactor() {
	// light child height greater than right child
	case 2:
		// LR first step
		// left rotate of left child
		//    |            |
		//    A            A
		//   /     =>     /
		//  B            C
		//   \          /
		//    C        B
		if n.left.getFactor() == -1 {
			n.left = n.left.leftRotate()
		}
		// right rotate
		//     |          |
		//     A          C
		//    /    =>    / \
		//   C          B   A
		//  /
		// B
		return n.rightRotate()
	// right child height greater than left child
	case -2:
		// RL first step
		// right rotate of right child
		//  |          |
		//  A          A
		//   \    =>    \
		//    B          C
		//   /            \
		//  C              B
		if n.right.getFactor() == 1 {
			n.right = n.right.rightRotate()
		}
		// left rotate
		//  |            |
		//  A            C
		//   \     =>   / \
		//    C        A   B
		//     \
		//      B
		return n.leftRotate()
	}
	n.maintain()
	return n
}

// leftRotate
//
//	  |           |
//	  A           C
//	 / \    =>   / \
//	B   C       A   E
//	   / \     / \
//	  D   E   B   D
func (n *_AvlTreeNode[K, V]) leftRotate() *_AvlTreeNode[K, V] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n
	n.maintain()
	newRoot.maintain()
	return newRoot
}

// rightRotate
//
//	    |           |
//	    A           B
//	   / \    =>   / \
//	  B   C       D   A
//	 / \             / \
//	D   E           E   C
func (n *_AvlTreeNode[K, V]) rightRotate() *_AvlTreeNode[K, V] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n
	n.maintain()
	newRoot.maintain()
	return newRoot
}

// maintain updates the node's height and size based on its children
func (n *_AvlTreeNode[K, V]) maintain() {
	n.height = 1
	n.size = 1
	if n.left != nil {
		n.height = max(n.height, n.left.height+1)
		n.size += n.left.size
	}
	if n.right != nil {
		n.height = max(n.height, n.right.height+1)
		n.size += n.right.size
	}
}

func (n *_AvlTreeNode[K, V]) inorder(yield func(K, V) bool) bool {
	if n.left != nil && !n.left.inorder(yield) {
		return false
	}
	if !yield(n.key, n.val) {
		return false
	}
	if n.right != nil && !n.right.inorder(yield) {
		return false
	}
	return true
}

// AvlTree represents a self-balancing binary search tree using AVL algorithm
type AvlTree[K, V any] struct {
	root *_AvlTreeNode[K, V]
	cmp  func(a, b K) int
}

// NewAvlTree creates an empty AVL tree with a given comparison function
func NewAvlTree[K, V any](cmp func(a, b K) int) *AvlTree[K, V] {
	return &AvlTree[K, V]{
		root: nil,
		cmp:  cmp,
	}
}

// Size returns the total number of elements.
// Complexity: O(1)
func (avl *AvlTree[K, V]) Size() int {
	if avl.root == nil {
		return 0
	}
	return avl.root.size
}

// Get searches for a value and returns (value, true) if found
// Complexity: O(log(n))
func (avl *AvlTree[K, V]) Get(key K) (val V, ok bool) {
	for node := avl.root; node != nil; {
		c := avl.cmp(key, node.key)
		if c == 0 {
			return node.val, true
		}
		if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return val, false
}

// Contains returns existence of a value
// Complexity: O(log(n))
func (avl *AvlTree[K, V]) Contains(key K) bool {
	_, ok := avl.Get(key)
	return ok
}

// Rank finds the k-th smallest element (1-based index)
// Complexity: O(log(n))
func (avl *AvlTree[K, V]) Rank(k int) (val V, ok bool) {
	if k <= 0 || k > avl.Size() {
		return val, false
	}
	for n := avl.root; ; {
		leftSize := 0
		if n.left != nil {
			leftSize = n.left.size
		}
		if k == leftSize+1 {
			return n.val, true
		}
		if k <= leftSize {
			n = n.left
			continue
		}
		n = n.right
		k -= leftSize + 1
	}
}

// Set inserts a value or updates if exists through compare function
// Complexity: O(log(n))
func (avl *AvlTree[K, V]) Set(key K, val V) { avl.root = avl.set(avl.root, key, val) }

func (avl *AvlTree[K, V]) set(n *_AvlTreeNode[K, V], key K, val V) *_AvlTreeNode[K, V] {
	if n == nil {
		return newNode(key, val)
	}
	c := avl.cmp(key, n.key)
	// renew value when equal
	if c == 0 {
		n.val = val
		return n
	}
	if c < 0 {
		n.left = avl.set(n.left, key, val)
	} else {
		n.right = avl.set(n.right, key, val)
	}
	return n.reBalance()
}

// Remove deletes a value from the tree if exists
// Complexity: O(log(n))
func (avl *AvlTree[K, V]) Remove(key K) { avl.root = avl.remove(avl.root, key) }

func (avl *AvlTree[K, V]) remove(n *_AvlTreeNode[K, V], key K) *_AvlTreeNode[K, V] {
	if n == nil {
		return nil
	}
	c := avl.cmp(key, n.key)
	if c < 0 {
		n.left = avl.remove(n.left, key)
	} else if c > 0 {
		n.right = avl.remove(n.right, key)
	} else {
		//  |       |
		//  A   =>  B
		//   \
		//    B
		if n.left == nil {
			return n.right
		}
		//    |       |
		//    A   =>  B
		//   /
		//  B
		if n.right == nil {
			return n.left
		}
		// find the successor of n, which is the leftmost node in n's right subtree
		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}
		// swap the value of n and successor
		n.key, successor.key = successor.key, n.key
		n.val, successor.val = successor.val, n.val
		// remove the successor
		n.right = avl.remove(n.right, successor.key)
	}
	return n.reBalance()
}

// Iter provides an in-order traversal iterator
func (avl *AvlTree[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if avl.root == nil {
			return
		}
		avl.root.inorder(yield)
	}
}
