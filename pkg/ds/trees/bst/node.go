package bst

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

type node[T ds.Numeric] struct {
	Left  *node[T]
	Right *node[T]
	Data  T
}

func newNode[T ds.Numeric](data T) *node[T] {
	return &node[T]{
		Left:  nil,
		Right: nil,
		Data:  data,
	}
}
