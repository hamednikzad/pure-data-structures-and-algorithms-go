package insertionsort

import "github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds"

func SortAscending[T ds.Numeric](array []T) []T {
	for i := 1; i <= len(array)-1; i++ {
		temp := array[i]
		j := i
		for j >= 1 && array[j-1]-temp > 0 {
			array[j] = array[j-1]
			j--
		}

		array[j] = temp
	}

	return array
}

func SortDescending[T ds.Numeric](array []T) []T {
	for i := 1; i <= len(array)-1; i++ {
		temp := array[i]
		j := i
		for j >= 1 && temp-array[j-1] > 0 {
			array[j] = array[j-1]
			j--
		}

		array[j] = temp
	}

	return array
}
