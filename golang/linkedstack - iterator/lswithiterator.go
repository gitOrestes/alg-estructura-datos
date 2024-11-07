package main

import (
	"errors"
	"fmt"
)

// interfaces
type LinkedStackInterface[T any] interface {
	Push(item T)
	Pop() (T, error)
	IsEmpty() bool
	Size() int
	NewIterator() *Iterator[T]
}

type IteratorInterface[T any] interface {
	Next() (T, bool)
	Position() int
	Add(item T)
	Remove() (T, error)
}

// struct
type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedStack[T any] struct {
	top  *Node[T]
	size int
}

type Iterator[T any] struct {
	current *Node[T]
	before  *Node[T]
	index   int
}

// constructores
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

func (ls *LinkedStack[T]) NewIterator() *Iterator[T] {
	return &Iterator[T]{
		current: ls.top,
		before:  nil,
		index:   0,
	}
}

// métodos de la linkedlist
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

// Métodos del iterador
func (i *Iterator[T]) Next() (T, bool) {
	if i.current == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := i.current.value
	i.before = i.current
	i.current = i.current.next
	i.index++
	return value, true
}

func (i *Iterator[T]) Position() int {
	return i.index
}

func (i *Iterator[T]) Add(item T) {
	newNode := NewNode(item)
	if i.current != nil {
		newNode.next = i.current.next
	}
	i.current.next = newNode
	i.Next()
}

func (i *Iterator[T]) Remove() (T, error) {
	if i.current == nil {
		return *new(T), fmt.Errorf("la pila esta vacia")
	}
	value := i.current.value
	i.before.next = i.current.next
	i.Next()

	return value, nil
}

func main() {
	stack := NewLinkedStack[int]()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	iterator := stack.NewIterator()

	// Usar el iterador para recorrer la pila
	fmt.Println("Elementos en la pila:")
	for {
		value, ok := iterator.Next()
		if !ok {
			break
		}
		fmt.Println(value)
	}

	// agregar un item
	iterator = stack.NewIterator()
	iterator.Next()
	iterator.Add(25)

	// ver la pila después de agregar
	fmt.Println("Pila después de agregar 25:")

	iterator = stack.NewIterator()
	for {
		value, ok := iterator.Next()
		if !ok {
			break
		}
		fmt.Println(value)
	}

	// Usar el iterador para eliminar un elemento
	iterator = stack.NewIterator()
	iterator.Next()
	removedValue, err := iterator.Remove()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Elemento eliminado: %d\n", removedValue)
	}

	// ver la pila después de eliminar
	fmt.Println("Pila después de eliminar el primer elemento:")
	iterator = stack.NewIterator()
	for {
		value, ok := iterator.Next()
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
