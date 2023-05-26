package avl

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func rotateLl[T ds.Numeric](parent *node[T]) *node[T] {
	pivot := parent.Left
	parent.Left = pivot.Right
	pivot.Right = parent
	return pivot
}

func rotateLr[T ds.Numeric](parent *node[T]) *node[T] {
	pivot := parent.Left
	parent.Left = rotateRr(pivot)
	return rotateLl(parent)
}

func rotateRl[T ds.Numeric](parent *node[T]) *node[T] {
	pivot := parent.Right
	parent.Right = rotateLl(pivot)
	return rotateRr(parent)
}

func rotateRr[T ds.Numeric](parent *node[T]) *node[T] {
	pivot := parent.Right
	parent.Right = pivot.Left
	pivot.Left = parent
	return pivot
}
