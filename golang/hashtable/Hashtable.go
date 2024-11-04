package main

import (
	"hash/fnv"
)

type Node[T any] struct {
	key   string
	value T
	next  *Node[T]
}

type HashTable[T any] struct {
	bucket []*Node[T]
	size   int
}

func hash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func NewHashTable[T any](size int) *HashTable[T] {
	return &HashTable[T]{
		bucket: make([]*Node[T], size),
		size:   size,
	}
}

func (ht *HashTable[T]) Insert(key string, value T) {
	index := hash(key) % uint64(ht.size)
	node := &Node[T]{key: key, value: value}

	if ht.bucket[index] == nil {
		ht.bucket[index] == node
	} else {
	}
}
