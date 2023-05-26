package avl

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

type Avl[T ds.Numeric] struct {
	Root *node[T]
}

func NewAvl[T ds.Numeric]() *Avl[T] {
	return &Avl[T]{Root: nil}
}

func (b *Avl[T]) Add(data T) bool {
	n := newNode(data)
	if b.Root == nil {
		b.Root = n
		return true
	}

	var success bool
	b.Root, success = insert(b.Root, n)

	return success
}

func insert[T ds.Numeric](current *node[T], newNode *node[T]) (*node[T], bool) {
	if current == nil {
		current = newNode
		return current, true
	}

	var success bool
	if newNode.Data < current.Data {
		current.Left, success = insert(current.Left, newNode)
		current = getBalancedTree(current)
	} else if newNode.Data > current.Data {
		current.Right, success = insert(current.Right, newNode)
		current = getBalancedTree(current)
	} else {
		success = false
	}

	return current, success
}

func getBalancedTree[T ds.Numeric](current *node[T]) *node[T] {
	bf := getBalanceFactor(current)

	if bf > 1 {
		if getBalanceFactor(current.Left) > 0 {
			return rotateLl(current)
		} else {
			return rotateLr(current)
		}
	}
	if bf < -1 {
		if getBalanceFactor(current.Right) > 0 {
			return rotateRl(current)
		} else {
			return rotateRr(current)
		}
	}

	return current
}

func getBalanceFactor[T ds.Numeric](node *node[T]) int {
	if node == nil {
		return 0
	}

	return getHeight(node.Left) - getHeight(node.Right)
}

func getHeight[T ds.Numeric](node *node[T]) int {
	if node == nil {
		return 0
	}

	var left = getHeight(node.Left)
	var right = getHeight(node.Right)
	return maxInt(left, right) + 1
}
