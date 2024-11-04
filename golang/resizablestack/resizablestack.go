package main

import (
	"errors"
	"fmt"
)

const initialStackCapacity = 3

type StackInterface[T any] interface {
	Push(item T)
	Pop() (T, error)
	IsEmpty() bool
	Size() int
	Capacity() int
}

type Stack[T any] struct {
	items    []T
	tail     int
	capacity int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items:    make([]T, initialStackCapacity),
		tail:     0,
		capacity: initialStackCapacity,
	}
}

func (s *Stack[T]) Push(item T) {
	if s.tail == s.capacity {
		s.Resize(s.capacity * 2)
	}
	s.items[s.tail] = item
	s.tail++
}

// agrega un elemento a la pila y redimensiona al doble cuando alcanza la capacidad maxima del arreglo
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, errors.New("la pila esta vacia")
	}
	item := s.items[s.tail-1]
	s.tail--

	//redimensionar si es necesario
	if s.tail == s.capacity/4 && s.capacity > initialStackCapacity {
		s.Resize(s.capacity / 2)
	}
	return item, nil
}

func (s *Stack[T]) Resize(newCapacity int) {
	newItems := make([]T, newCapacity)
	for i := 0; i < s.tail; i++ {
		newItems[i] = s.items[i]
	}
	s.items = newItems
	s.capacity = newCapacity
}

func (s *Stack[T]) IsEmpty() bool {
	return s.tail == 0
}

func (s *Stack[T]) Size() int {
	return s.tail
}

func (s *Stack[T]) Capacity() int {
	return s.capacity
}

func main() {
	var stack StackInterface[int] = NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	fmt.Println("capacidad original de la cola:", stack.Capacity())

	for !stack.IsEmpty() {
		item, err := stack.Pop()
		if err == nil {
			fmt.Println("popped: ", item)
		}
	}

	fmt.Println("capacidad original de la cola:", stack.Capacity())
}
