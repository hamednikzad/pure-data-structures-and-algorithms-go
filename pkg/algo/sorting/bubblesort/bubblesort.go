package bubblesort

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func SortAscending[T ds.Numeric](array []T) []T {
	swapped := true
	for i := len(array) - 1; i >= 0 && swapped; i-- {
		swapped = false
		for j := 0; j <= i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
	}

	return array
}

func SortDescending[T ds.Numeric](array []T) []T {
	swapped := true
	for i := len(array) - 1; i >= 0 && swapped; i-- {
		swapped = false
		for j := 0; j <= i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
	}

	return array
}
