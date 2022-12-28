package genericlist

import (
	"errors"
	"fmt"
)

type List[T comparable] struct {
	items    []T
	Count    int
	Capacity int
}

func New[T comparable](capacity int) *List[T] {
	list := &List[T]{}
	if capacity > 0 {
		list.items = make([]T, capacity)
	}
	return list
}

func (list *List[T]) Print() {
	fmt.Printf("Capacity:%d, Count:%d\t", list.Capacity, list.Count)
	for i := 0; i < list.Count; i++ {
		sep := ", "
		if i == list.Count-1 {
			sep = ""
		}
		fmt.Print(list.items[i], sep)
	}
	fmt.Println()
}

func (list *List[T]) GetValue(index int) T {
	if index < 0 || index >= list.Count {
		return *new(T)
	}
	return list.items[index]
}

func (list *List[T]) Add(value T) {
	if list.Count == len(list.items) {
		list.checkCapacity(list.Count + 1)
	}

	list.items[list.Count] = value
	list.Count++
}

func (list *List[T]) checkCapacity(size int) {
	if len(list.items) >= size {
		return
	}

	var newCapacity int
	if len(list.items) == 0 {
		newCapacity = 4
	} else {
		newCapacity = list.Count * 2
	}

	if newCapacity < size {
		newCapacity = size
	}

	list.changeCapacity(newCapacity)
}

func (list *List[T]) changeCapacity(newCapacity int) {
	if newCapacity <= list.Count {
		panic("Capacity out of range")
	}

	if newCapacity == list.Capacity {
		return
	}

	if newCapacity > 0 {
		newItems := make([]T, newCapacity)
		if list.Count > 0 {
			for i := 0; i < list.Count; i++ {
				newItems[i] = list.items[i]
			}
		}
		list.items = newItems
	} else {
		list.items = make([]T, 0)
	}

	list.Capacity = newCapacity
}

func (list *List[T]) Clear() {
	if list.Count <= 0 {
		return
	}
	for i := 0; i < list.Count; i++ {
		list.items[i] = *new(T)
	}
	list.Count = 0
}

func (list *List[T]) Contains(item T) bool {
	for i := 0; i < list.Count; i++ {
		if list.items[i] == item {
			return true
		}
	}
	return false
	/*if item == nil {
		for i := 0; i < list.Count; i++ {
			if list.items[i] == nil {
				return true
			}
		}
		return false
	}else{
		for i := 0; i < list.Count; i++ {
			if list.items[i] == item {
				return true
			}
		}
		return false
	}*/
}

func (list *List[T]) IndexOf(item T) int {
	for i := 0; i < list.Count; i++ {
		if list.items[i] == item {
			return i
		}
	}
	return -1
}

func copyArray[T comparable](array []T, srcIndex, dstIndex int) {
	for i := dstIndex; i > srcIndex; i-- {
		array[i] = array[i-1]
	}
}

func (list *List[T]) Insert(index int, item T) error {
	if index < 0 || index > list.Count {
		return errors.New("index out of range")
	}
	if list.Count == len(list.items) {
		list.checkCapacity(list.Count + 1)
	}
	if index < list.Count {
		copyArray(list.items, index, list.Count)
	}
	list.items[index] = item
	list.Count++
	return nil
}
