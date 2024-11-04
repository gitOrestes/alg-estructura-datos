package main

import (
	"errors"
	"fmt"
)

type LinkedQueueInterface[T any] interface {
	Push(item T)
	Pop() (T, error)
	IsEmpty() bool
	Size() int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedQueue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		next:  nil,
	}
}

func (lq *LinkedQueue[T]) Push(item T) {
	newNode := NewNode(item)

	if lq.IsEmpty() {
		lq.head = newNode
	}
	if !lq.IsEmpty() {
		lq.tail.next = newNode
	}
	lq.tail = newNode
	lq.size++
}

func (lq *LinkedQueue[T]) Pop() (T, error) {
	if lq.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("la fila esta vacia")
	}
	item := lq.head.value
	if lq.size > 1 {
		lq.head = lq.head.next
	} else {
		lq.head = nil
	}
	lq.size--

	return item, nil
}

func (lq *LinkedQueue[T]) IsEmpty() bool {
	return lq.size == 0
}

func (lq *LinkedQueue[T]) Size() int {
	return lq.size
}

func main() {
	var lq LinkedQueueInterface[int] = NewLinkedQueue[int]()
	lq.Push(1)
	lq.Push(2)
	lq.Push(3)
	lq.Push(4)

	for !lq.IsEmpty() {
		item, _ := lq.Pop()
		fmt.Println(item)
	}

	item, err := lq.Pop()
	if err == nil {
		fmt.Println(item)
	} else {
		fmt.Println("error:", err)
	}
}
