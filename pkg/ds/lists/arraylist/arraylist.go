package arraylist

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	items    []interface{}
	Count    int
	Capacity int
}

func New(capacity int) *ArrayList {
	aList := &ArrayList{}
	if capacity > 0 {
		aList.items = make([]interface{}, capacity)
	}
	return aList
}

func (list *ArrayList) Print() {
	fmt.Printf("Capacity:%d, Count:%d\t", list.Capacity, list.Count)
	for i := 0; i < list.Count; i++ {
		fmt.Print(list.items[i], " ")
	}
	fmt.Println()
}

func (list *ArrayList) GetValue(index int) interface{} {
	if index < 0 || index >= list.Count {
		return nil
	}
	return list.items[index]
}

func (list *ArrayList) Add(value interface{}) {
	if list.Count == len(list.items) {
		list.checkCapacity(list.Count + 1)
	}

	list.items[list.Count] = value
	list.Count++
}

func (list *ArrayList) checkCapacity(size int) {
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

func (list *ArrayList) changeCapacity(newCapacity int) {
	if newCapacity <= list.Count {
		panic("Capacity out of range")
	}

	if newCapacity == list.Capacity {
		return
	}

	if newCapacity > 0 {
		newItems := make([]interface{}, newCapacity)
		if list.Count > 0 {
			for i := 0; i < list.Count; i++ {
				newItems[i] = list.items[i]
			}
		}
		list.items = newItems
	} else {
		list.items = make([]interface{}, 0)
	}

	list.Capacity = newCapacity
}

func (list *ArrayList) Clear() {
	if list.Count <= 0 {
		return
	}
	for i := 0; i < list.Count; i++ {
		list.items[i] = nil
	}
	list.Count = 0
}

func (list *ArrayList) Contains(item interface{}) bool {
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

func (list *ArrayList) IndexOf(item interface{}) int {
	for i := 0; i < list.Count; i++ {
		if list.items[i] == item {
			return i
		}
	}
	return -1
}

func copyArray(array []interface{}, srcIndex, dstIndex int) {
	for i := dstIndex; i > srcIndex; i-- {
		array[i] = array[i-1]
	}
}

func (list *ArrayList) Insert(index int, item interface{}) error {
	if index < 0 || index > list.Count {
		return errors.New("index out of range")
	}
	if list.Count == len(list.items) {
		list.checkCapacity(list.Count + 1)
	}
	if index < list.Count {
		copyArray(list.items, index, list.Count)
		//for i := list.Count; i > index; i-- {
		//	list.items[i] = list.items[i-1]
		//}
	}
	list.items[index] = item
	list.Count++
	return nil
}
