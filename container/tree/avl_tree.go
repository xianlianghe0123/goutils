package tree

import (
	"iter"
)

type _AvlTreeNode[T any] struct {
	val T
	// height means the max height of subtree that this node is the root
	height int
	// size is the subtree's node count
	size        int
	left, right *_AvlTreeNode[T]
}

func newNode[T any](val T) *_AvlTreeNode[T] {
	return &_AvlTreeNode[T]{
		val:    val,
		height: 1,
		size:   1,
	}
}

func (n *_AvlTreeNode[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// getFactor returns the difference between left child height and right child height
func (n *_AvlTreeNode[T]) getFactor() int { return n.left.getHeight() - n.right.getHeight() }

func (n *_AvlTreeNode[T]) reBalance() *_AvlTreeNode[T] {
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
func (n *_AvlTreeNode[T]) leftRotate() *_AvlTreeNode[T] {
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
func (n *_AvlTreeNode[T]) rightRotate() *_AvlTreeNode[T] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n
	n.maintain()
	newRoot.maintain()
	return newRoot
}

func (n *_AvlTreeNode[T]) maintain() {
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

func (n *_AvlTreeNode[T]) inorder(yield func(T) bool) bool {
	if n.left != nil && !n.left.inorder(yield) {
		return false
	}
	if !yield(n.val) {
		return false
	}
	if n.right != nil && !n.right.inorder(yield) {
		return false
	}
	return true
}

type AvlTree[T any] struct {
	root *_AvlTreeNode[T]
	cmp  func(a, b T) int
}

func NewAvlTree[T any](cmp func(a, b T) int) *AvlTree[T] {
	return &AvlTree[T]{
		root: nil,
		cmp:  cmp,
	}
}

func (avl *AvlTree[T]) Size() int {
	if avl.root == nil {
		return 0
	}
	return avl.root.size
}

func (avl *AvlTree[T]) Get(val T) (ret T, ok bool) {
	for node := avl.root; node != nil; {
		c := avl.cmp(val, node.val)
		if c == 0 {
			return node.val, true
		}
		if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return ret, false
}

func (avl *AvlTree[T]) Contains(val T) bool {
	_, ok := avl.Get(val)
	return ok
}

func (avl *AvlTree[T]) Rank(k int) (val T, ok bool) {
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

func (avl *AvlTree[T]) Set(val T) { avl.root = avl.set(avl.root, val) }

func (avl *AvlTree[T]) set(n *_AvlTreeNode[T], val T) *_AvlTreeNode[T] {
	if n == nil {
		return newNode(val)
	}
	c := avl.cmp(val, n.val)
	// renew value when equal
	if c == 0 {
		n.val = val
		return n
	}
	if c < 0 {
		n.left = avl.set(n.left, val)
	} else {
		n.right = avl.set(n.right, val)
	}
	return n.reBalance()
}

func (avl *AvlTree[T]) Remove(val T) { avl.root = avl.remove(avl.root, val) }

func (avl *AvlTree[T]) remove(n *_AvlTreeNode[T], val T) *_AvlTreeNode[T] {
	if n == nil {
		return nil
	}
	c := avl.cmp(val, n.val)
	if c < 0 {
		n.left = avl.remove(n.left, val)
	} else if c > 0 {
		n.right = avl.remove(n.right, val)
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
		n.val, successor.val = successor.val, n.val
		// remove the successor
		n.right = avl.remove(n.right, successor.val)
	}
	return n.reBalance()
}

func (avl *AvlTree[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		if avl.root == nil {
			return
		}
		avl.root.inorder(yield)
	}
}
