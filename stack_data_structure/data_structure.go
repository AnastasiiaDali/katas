package stack_data_structure

import (
	"log"
)

type Stack struct {
	stack       []int
	lastElement int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		stack:       make([]int, capacity),
		lastElement: -1,
	}
}

func (s *Stack) Push(item int) bool {
	if s.IsFull() {
		log.Print("stack is full, can not push new item")
		return false
	}
	s.lastElement += 1
	s.stack[s.lastElement] = item
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		log.Print("stack is empty, nothing to pop")
		return 0, false
	}
	popItem := s.stack[s.lastElement]
	s.stack = s.stack[:popItem]
	return popItem, true
}

func (s Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		log.Print("stack is empty, nothing to peek")
		return 0, false
	}
	peekedItem := s.stack[s.lastElement]
	return peekedItem, true
}

func (s *Stack) IsEmpty() bool {
	return s.lastElement == -1
}

func (s *Stack) IsFull() bool {
	return s.lastElement == len(s.stack)-1
}
