package stack

import "sync/atomic"

type Node[T any] struct {
	Value T
	Next *Node[T]
}

type TreiberStack[T any] struct {
	head atomic.Pointer[Node[T]]
}