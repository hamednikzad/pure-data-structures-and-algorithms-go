package doublylinkedlist

import "fmt"

type node[T any] struct {
	Data T
	Next *node[T]
	Prev *node[T]
}

type DoublyLinkedList[T comparable] struct {
	head  *node[T]
	Count int
}

func New[T comparable]() *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{}
	return list
}

func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.Count == 0
}

func (list *DoublyLinkedList[T]) AddFirst(value T) {
	newNode := &node[T]{
		Data: value,
		Next: nil,
		Prev: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.Count++
		return
	}

	newNode.Next = list.head
	list.head.Prev = newNode
	list.head = newNode
	list.Count++
}

func (list *DoublyLinkedList[T]) AddLast(value T) {
	newNode := &node[T]{
		Data: value,
		Next: nil,
		Prev: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.Count++
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}

	newNode.Prev = current
	current.Next = newNode
	list.Count++
}

func (list *DoublyLinkedList[T]) find(value T) *node[T] {
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

func (list *DoublyLinkedList[T]) Contains(value T) bool {
	return list.find(value) != nil
}

func (list *DoublyLinkedList[T]) Remove(index int) {
	if list.head == nil || index < 0 || index >= list.Count {
		return
	}
	if index == 0 {
		list.head = list.head.Next

		if list.head != nil {
			list.head.Prev = nil
		}
	} else {
		current := list.head
		prev := list.head
		var i = 0
		for current != nil && i < index {
			prev = current
			current = current.Next
			i++
		}

		prev.Next = current.Next

		if current.Next != nil {
			current.Next.Prev = prev
		}
	}

	list.Count--
}

func (list *DoublyLinkedList[T]) RemoveLast() {
	list.Remove(list.Count - 1)
}

func (list *DoublyLinkedList[T]) RemoveFirst() {
	list.Remove(0)
}

func (list *DoublyLinkedList[T]) RemoveValue(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.Data == value {
		list.head = list.head.Next

		if list.head != nil {
			list.head.Prev = nil
		}
		list.Count--
		return true
	}

	current := list.head
	prev := list.head
	for current != nil {
		if current.Data == value {
			prev.Next = current.Next

			if current.Next != nil {
				current.Next.Prev = prev
			}

			list.Count--
			return true
		}

		prev = current
		current = current.Next
	}

	return false
}

func (list *DoublyLinkedList[T]) GetValueAtIndex(index int) T {
	if list.head == nil || index < 0 || index >= list.Count {
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

func (list *DoublyLinkedList[T]) Traverse(action func(a ...any) (n int, err error)) {
	if list.head == nil {
		return
	}
	current := list.head
	for current != nil {
		_, _ = action(current.Data)
		current = current.Next
	}
}

func (list *DoublyLinkedList[T]) Print() {
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
