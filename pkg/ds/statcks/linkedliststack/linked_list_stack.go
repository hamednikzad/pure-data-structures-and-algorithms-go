package linkedliststack

import (
	"errors"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/singlylinkedlist"
)

type LinkedListStack[T comparable] struct {
	list *singlylinkedlist.SinglyLinkedList[T]
}

func New[T comparable]() *LinkedListStack[T] {
	queue := &LinkedListStack[T]{}
	queue.list = singlylinkedlist.New[T]()
	return queue
}

func (stack *LinkedListStack[T]) IsEmpty() bool {
	return stack.list.IsEmpty()
}

func (stack *LinkedListStack[T]) GetCount() int {
	return stack.list.Count
}

func (stack *LinkedListStack[T]) Push(value T) {
	stack.list.AddFirst(value)
}

func (stack *LinkedListStack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		return *new(T), errors.New("the Stack is empty")
	}

	value := stack.list.GetValueAtIndex(0)
	stack.list.RemoveFirst()

	return value, nil
}

func (stack *LinkedListStack[T]) Peek() (T, error) {
	if stack.IsEmpty() {
		return *new(T), errors.New("the Stack is empty")
	}

	value := stack.list.GetValueAtIndex(0)

	return value, nil
}

func (stack *LinkedListStack[T]) Traverse(action func(a ...any) (n int, err error)) {
	stack.list.Traverse(action)
}

func (stack *LinkedListStack[T]) Print() {
	stack.list.Print()
}
