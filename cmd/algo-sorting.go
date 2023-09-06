package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/algo/sorting/bubblesort"
)

func bubbleSort() {
	array := []int{2, 6, 1, 8, 5}
	ascending := bubblesort.SortAscending(array)
	fmt.Print("Ascending: ")
	for _, t := range ascending {
		fmt.Printf("%d ", t)
	}
	fmt.Println()

	descending := bubblesort.SortDescending(array)
	fmt.Print("Descending: ")
	for _, t := range descending {
		fmt.Printf("%d ", t)
	}
}
