package shellsort

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func SortAscending[T ds.Numeric](array []T) []T {
	size := len(array)
	for interval := size / 2; interval > 0; interval /= 2 {
		for i := interval; i < size; i++ {
			currentKey := array[i]
			k := i
			for k >= interval && array[k-interval] > currentKey {
				array[k] = array[k-interval]
				k -= interval
			}
			array[k] = currentKey
		}
	}

	return array
}

func SortDescending[T ds.Numeric](array []T) []T {
	size := len(array)
	for interval := size / 2; interval > 0; interval /= 2 {
		for i := interval; i < size; i++ {
			currentKey := array[i]
			k := i
			for k >= interval && array[k-interval] < currentKey {
				array[k] = array[k-interval]
				k -= interval
			}
			array[k] = currentKey
		}
	}

	return array
}
