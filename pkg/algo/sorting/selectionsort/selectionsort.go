package selectionsort

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func SortAscending[T ds.Numeric](array []T) []T {
	for i := 0; i <= len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}

	return array
}

func SortDescending[T ds.Numeric](array []T) []T {
	for i := 0; i <= len(array)-1; i++ {
		max := i
		for j := i + 1; j < len(array); j++ {
			if array[j] > array[max] {
				max = j
			}
		}

		if max != i {
			array[i], array[max] = array[max], array[i]
		}
	}

	return array
}
