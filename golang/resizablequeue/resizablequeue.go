package main

import (
	"errors"
	"fmt"
)

const initialCapacity = 3

type QueueInterface[T any] interface {
	Push(item T)
	Pop() (T, error)
	IsEmpty() bool
	Size() int
	Capacity() int
}

type Queue[T any] struct {
	items    []T
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items:    make([]T, initialCapacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: initialCapacity,
	}
}

// agrega un elemento a la cola y redimensiona al doble cuando alcanza la capacidad maxima del arreglo
func (q *Queue[T]) Push(item T) {
	if q.size == q.capacity {
		q.resize(q.capacity * 2)
	}
	q.items[q.tail] = item
	q.tail++
	q.size++
}

func (q *Queue[T]) Pop() (T, error) {
	if q.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("la cola esta vacia")
	}
	item := q.items[q.head]
	q.head++
	q.size--

	// Redimensionar si es necesario
	if q.size <= q.capacity/4 && q.capacity > initialCapacity {
		q.resize(q.capacity / 2)
	}

	return item, nil
}

func (q *Queue[T]) resize(newCapacity int) {
	newItems := make([]T, newCapacity)
	for i := 0; i < q.size; i++ {
		newItems[i] = q.items[q.head+i]
	}
	q.items = newItems
	q.head = 0
	q.tail = q.size
	q.capacity = newCapacity
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Capacity() int {
	return q.capacity
}

func main() {

	var queue QueueInterface[int] = NewQueue[int]()
	fmt.Println("capacidad original de la cola:", queue.Capacity())
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)

	fmt.Println("tamaño después de pushear:", queue.Capacity())

	for !queue.IsEmpty() {
		item, err := queue.Pop()
		if err == nil {
			fmt.Println("Popped:", item)
		}
	}

	fmt.Println("tamaño después de pop:", queue.Capacity())
}
