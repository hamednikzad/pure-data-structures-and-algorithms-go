package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/algo/sorting/bubblesort"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/algo/sorting/insertionsort"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/algo/sorting/selectionsort"
)

func print(title string, arr []int) {
	fmt.Printf("%s: ", title)
	for _, t := range arr {
		fmt.Printf("%d ", t)
	}
	fmt.Println()
}

func makeArray() []int {
	arr := []int{2, 6, 1, 8, 5}
	print("Original", arr)
	return arr
}

func bubbleSort() {
	fmt.Println("bubbleSort")
	array := makeArray()
	ascending := bubblesort.SortAscending(array)
	print("Ascending", ascending)

	descending := bubblesort.SortDescending(array)
	print("Descending", descending)
}

func insertionSort() {
	fmt.Println("insertionSort")
	array := makeArray()
	ascending := insertionsort.SortAscending(array)
	print("Ascending", ascending)

	descending := insertionsort.SortDescending(array)
	print("Descending", descending)
}

func selectionSort() {
	fmt.Println("selectionSort")
	array := makeArray()
	ascending := selectionsort.SortAscending(array)
	print("Ascending", ascending)

	descending := selectionsort.SortDescending(array)
	print("Descending", descending)
}
