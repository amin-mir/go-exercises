package bt

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	val         T
	left, right *Node[T]
}

func New[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{val: val}
}

func (n *Node[T]) Insert(val T) {
	if val == n.val {
		return
	}

	if val < n.val {
		if n.left != nil {
			n.left.Insert(val)
			return
		}
		n.left = &Node[T]{val: val}
		return
	}

	if n.right != nil {
		n.right.Insert(val)
		return
	}
	n.right = &Node[T]{val: val}
}

func (n *Node[T]) Search(val T) bool {
	if n == nil {
		return false
	}

	if n.val == val {
		return true
	}

	if val < n.val {
		return n.left.Search(val)
	}

	return n.right.Search(val)
}

func (n *Node[T]) Sorted() (res []T) {
	return n.appendSorted(res)
}

func (n *Node[T]) appendSorted(sorted []T) []T {
	if n.left != nil {
		sorted = n.left.appendSorted(sorted)
	}
	sorted = append(sorted, n.val)
	if n.right != nil {
		sorted = n.right.appendSorted(sorted)
	}
	return sorted
}

func (n *Node[T]) Delete(val T) {
	if n.val == val {
		switch n.numChildren() {
		case 1:
			if n.left != nil {
				n.val = n.left.val
				n.left = nil
			} else {
				n.val = n.right.val
				n.right = nil
			}
		case 2:
			// Find the minimum in the right subtree.
			minNode := n.right.deleteMin(n)
			n.val = minNode.val
		default:
			panic("should not come here...")
		}
		return
	}

	// Value in the left subtree.
	if val < n.val {
		if n.left.val == val && n.left.isLeaf() {
			n.left = nil
			return
		}
		n.left.Delete(val)
		return
	}

	// Value in the right subtree.
	if n.right.val == val && n.right.isLeaf() {
		n.right = nil
		return
	}
	n.right.Delete(val)
}

func (n *Node[T]) deleteMin(parent *Node[T]) *Node[T] {
	if n.left != nil {
		return n.left.deleteMin(n)
	}
	parent.left = nil
	return n
}

func (n *Node[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node[T]) numChildren() (num int) {
	if n.left != nil {
		num++
	}
	if n.right != nil {
		num++
	}
	return
}

func (n *Node[T]) Print() {
	print([]*Node[T]{n})
}

func print[T constraints.Ordered](nodes []*Node[T]) {
	children := make([]*Node[T], 0, len(nodes)*2)
	for _, n := range nodes {
		fmt.Printf(" %v ", n.val)
		if n.left != nil {
			children = append(children, n.left)
		}
		if n.right != nil {
			children = append(children, n.right)
		}
	}
	fmt.Println()

	if len(children) > 0 {
		print(children)
	}
}

func (n *Node[T]) String() string {
	left := "null"
	if n.left != nil {
		left = fmt.Sprintf("%v", n.left.val)
	}

	right := "null"
	if n.right != nil {
		right = fmt.Sprintf("%v", n.right.val)
	}

	return fmt.Sprintf("%v => (%v, %v)", n.val, left, right)
}
