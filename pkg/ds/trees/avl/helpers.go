package avl

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func maxInt[T ds.Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}
