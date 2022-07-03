package stack_data_structure

import (
	"log"
	"strconv"
)

type Stack struct {
	stack       []string
	lastElement int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		stack:       make([]string, capacity),
		lastElement: -1,
	}
}

func (s *Stack) Push(item string) bool {
	if s.IsFull() {
		log.Print("stack is full, can not push new item")
		return false
	}
	s.lastElement += 1
	s.stack[s.lastElement] = item
	return true
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		log.Print("stack is empty, nothing to pop")
		return "zero", false
	}
	popItem := s.stack[s.lastElement]
	i, _ := strconv.Atoi(popItem)
	s.stack = s.stack[:i]
	return popItem, true
}

func (s Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		log.Print("stack is empty, nothing to peek")
		return "zero", false
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
