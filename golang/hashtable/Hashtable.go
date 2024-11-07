package main

import (
	"fmt"
	"hash/fnv"
)

type Node[T any] struct {
	key   string
	value T
	next  *Node[T]
}

type Hashtable[T any] struct {
	bucket []*Node[T]
	size   int
	count  int
}

func hash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func NewHashtable[T any](size int) *Hashtable[T] {
	return &Hashtable[T]{
		bucket: make([]*Node[T], size),
		size:   size,
	}
}

func (ht *Hashtable[T]) Insert(key string, value T) {
	index := hash(key) % uint64(ht.size)

	if ht.bucket[index] == nil {
		ht.bucket[index] = &Node[T]{key: key, value: value}
	} else {
		current := ht.bucket[index]
		for current.next != nil {
			current = current.next
		}
		current.next = &Node[T]{key: key, value: value}
	}

	ht.count++

	if float64(ht.count)/float64(ht.size) >= 2.0 {
		ht.resize()
	}
}

func (ht *Hashtable[T]) Delete(key string) bool {
	index := hash(key) % uint64(ht.size)

	current := ht.bucket[index]
	var prev *Node[T]

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.bucket[index] = current.next
			} else {
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}

	return false
}

func (ht *Hashtable[T]) resize() {
	newSize := ht.size * 2
	newTable := NewHashtable[T](newSize)

	for i := 0; i < ht.size; i++ {
		current := ht.bucket[i]
		for current != nil {
			newTable.Insert(current.key, current.value)
			current = current.next
		}
	}

	ht.bucket = newTable.bucket
	ht.size = newSize

}

func (ht *Hashtable[T]) Print() {
	for i := 0; i < ht.size; i++ {
		fmt.Printf("Bucket %d: ", i)
		current := ht.bucket[i]
		for current != nil {
			fmt.Printf("(%s: %v) ", current.key, current.value)
			current = current.next
		}
		fmt.Println()
	}
}

func main() {
	// Crear una nueva tabla hash
	ht := NewHashtable[int](2)

	// Insertar algunos valores
	ht.Insert("apple", 1)
	ht.Insert("banana", 2)
	ht.Print()
	ht.Insert("cherry", 3)
	ht.Insert("date", 4)
	ht.Insert("elderberry", 5) // Esto debería activar el redimensionado

	// Imprimir la tabla antes del redimensionado
	fmt.Println("Antes de redimensionar:")
	ht.Print()

	// Eliminar un valor
	ht.Delete("banana")

	// Imprimir después de eliminar
	fmt.Println("\nDespués de eliminar 'banana':")
	ht.Print()
}
