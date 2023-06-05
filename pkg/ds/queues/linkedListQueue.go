package queues

import (
	"errors"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/ds/lists/singlylinkedlist"
)

type LinkedListQueue[T comparable] struct {
	list  *singlylinkedlist.SinglyLinkedList[T]
	Count int
}

func New[T comparable]() *LinkedListQueue[T] {
	queue := &LinkedListQueue[T]{}
	queue.list = singlylinkedlist.New[T]()
	queue.Count = 0
	return queue
}

func (queue *LinkedListQueue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *LinkedListQueue[T]) Enqueue(value T) {
	queue.list.AddLast(value)
}

func (queue *LinkedListQueue[T]) Dequeue() (T, error) {
	if queue.IsEmpty() {
		return *new(T), errors.New("the Queue is empty")
	}

	value := queue.list.GetValueAtIndex(0)
	queue.list.RemoveFirst()

	return value, nil
}

func (queue *LinkedListQueue[T]) Peek() (T, error) {
	if queue.IsEmpty() {
		return *new(T), errors.New("the Queue is empty")
	}

	value := queue.list.GetValueAtIndex(0)

	return value, nil
}

func (queue *LinkedListQueue[T]) Print() {
	queue.list.Print()
}
