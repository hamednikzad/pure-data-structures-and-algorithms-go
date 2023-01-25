package sortedlinkedlist

import (
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type node[T Numeric] struct {
	Data T
	Next *node[T]
}

type SortedLinkedList[T Numeric] struct {
	head  *node[T]
	Count int
}

func New[T Numeric]() *SortedLinkedList[T] {
	list := &SortedLinkedList[T]{}
	return list
}

func (list *SortedLinkedList[T]) IsEmpty() bool {
	return list.Count == 0
}

func (list *SortedLinkedList[T]) Add(value T) {
	newNode := &node[T]{
		Data: value,
		Next: nil,
	}

	if list.head == nil || value <= list.head.Data {
		newNode.Next = list.head
		list.head = newNode
		list.Count++
		return
	}

	current := list.head
	for current.Next != nil && current.Next.Data < newNode.Data {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	list.Count++
}

func (list *SortedLinkedList[T]) find(value T) *node[T] {
	if list.head == nil {
		return nil
	}

	current := list.head
	for current != nil {
		if current.Data == value {
			return current
		}

		current = current.Next
	}
	return nil
}

func (list *SortedLinkedList[T]) Contains(value T) bool {
	return list.find(value) != nil
}

func (list *SortedLinkedList[T]) RemoveValue(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.Data == value {
		list.head = list.head.Next

		list.Count--
		return true
	}

	current := list.head
	prev := list.head
	for current != nil {
		if current.Data == value {
			prev.Next = current.Next

			list.Count--
			return true
		}

		prev = current
		current = current.Next
	}

	return false
}

func (list *SortedLinkedList[T]) GetValueAtIndex(index int) T {
	if list.head == nil || index < 0 || index > list.Count {
		return *new(T)
	}

	i := 0
	current := list.head
	for i != index && current != nil {
		current = current.Next
		i++
	}

	return current.Data
}

func (list *SortedLinkedList[T]) Traverse(action func(a ...any) (n int, err error)) {
	if list.head == nil {
		return
	}
	current := list.head
	for current != nil {
		_, _ = action(current.Data)
		current = current.Next
	}
}

func (list *SortedLinkedList[T]) Print() {
	fmt.Printf("Count:%d =>\t", list.Count)
	current := list.head
	sep := ", "
	for current != nil {
		if current.Next == nil {
			sep = ""
		}
		fmt.Print(current.Data, sep)
		current = current.Next
	}
	fmt.Println()
}
