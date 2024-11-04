package main

import (
	"errors"
	"fmt"
)

type LinkedStackInterface[T any] interface {
	Push(item T)
	Pop() (T, error)
	IsEmpty() bool
	Size() int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedStack[T any] struct {
	top  *Node[T]
	size int
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{
		top:  nil,
		size: 0,
	}
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
		next:  nil,
	}
}

func (ls *LinkedStack[T]) Push(item T) {
	newNode := NewNode(item)
	newNode.next = ls.top
	ls.top = newNode
	ls.size++
}

func (ls *LinkedStack[T]) Pop() (T, error) {
	if ls.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("la pila esta vacia")
	}
	itemValue := ls.top.value
	ls.top = ls.top.next
	ls.size--

	return itemValue, nil
}

func (ls *LinkedStack[T]) IsEmpty() bool {
	return ls.size == 0
}

func (ls *LinkedStack[T]) Size() int {
	return ls.size
}

func main() {
	var linkedStack LinkedStackInterface[int] = NewLinkedStack[int]()
	linkedStack.Push(1)
	linkedStack.Push(2)
	linkedStack.Push(3)

	for !linkedStack.IsEmpty() {
		val, err := linkedStack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Popped value:", val)
		}
	}
}
