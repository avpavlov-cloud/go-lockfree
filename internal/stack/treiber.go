package stack

import "sync/atomic"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type TreiberStack[T any] struct {
	head atomic.Pointer[Node[T]]
}

func NewTreiberStack[T any]() *TreiberStack[T] {
	return &TreiberStack[T]{}
}

func (s *TreiberStack[T]) Push(val T) {
	newNode := &Node[T]{Value: val}

	for {
		oldHead := s.head.Load()
		newNode.Next = oldHead
		if s.head.CompareAndSwap(oldHead, newNode) {
			return
		}
	}
}

func (s *TreiberStack[T]) Pop() (T, bool) {
	for {
		oldHead := s.head.Load()

		if oldHead == nil {
			var zero T
			return zero, false
		}

		next := oldHead.Next

		if s.head.CompareAndSwap(oldHead, next) {
			return oldHead.Value, true
		}
	}
}
