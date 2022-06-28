package stack_data_structure

import "log"

type Item interface {
	int | string
}
type Stack[T Item] struct {
	stack       []T
	lastElement int
}

func NewStack[T Item](capacity int) *Stack[T] {
	return &Stack[T]{
		stack:       make([]T, capacity),
		lastElement: -1,
	}
}

func (s *Stack[T]) Push(item T) bool {
	if s.IsFull() {
		log.Print("stack is full, can not push new item")
		return false
	}
	s.lastElement += 1
	s.stack[s.lastElement] = item
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var nothing T
		log.Print("stack is empty, nothing to pop")
		return nothing, false
	}
	popItem := s.stack[s.lastElement]
	s.stack = s.stack[:len(s.stack)-1]
	return popItem, true
}

func (s Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		log.Print("stack is empty, nothing to peek")
		return 0, false
	}
	peekedItem := s.stack[s.lastElement]
	return peekedItem, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.lastElement == -1
}

func (s *Stack[T]) IsFull() bool {
	return s.lastElement == len(s.stack)-1
}
