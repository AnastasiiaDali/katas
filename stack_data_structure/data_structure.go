package stack_data_structure

import (
	"log"
)

type Stack struct {
	StackOfItems []string
	lastElement  int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		StackOfItems: make([]string, capacity),
		lastElement:  -1,
	}
}

func (s *Stack) Push(item string) bool {
	if s.IsFull() {
		log.Print("StackOfItems is full, can not push new item")
		return false
	}
	s.lastElement += 1
	s.StackOfItems[s.lastElement] = item
	return true
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		log.Print("StackOfItems is empty, nothing to pop")
		return "zero", false
	}
	popItem := s.StackOfItems[s.lastElement]
	s.lastElement--
	return popItem, true
}

func (s Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		log.Print("StackOfItems is empty, nothing to peek")
		return "zero", false
	}
	peekedItem := s.StackOfItems[s.lastElement]
	return peekedItem, true
}

func (s *Stack) IsEmpty() bool {
	return s.lastElement == -1
}

func (s *Stack) IsFull() bool {
	return s.lastElement == len(s.StackOfItems)-1
}
