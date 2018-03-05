package main

import "errors"

type stack struct {
	stack []int
}

func (s stack) getLength() int {
	return len(s.stack)
}

func newStack() stack {
	return stack{
		stack: []int{},
	}
}

func (s *stack) push(element int) {
	s.stack = append(s.stack, element)
}

func (s *stack) pop() (element int, err error) {
	if (*s).getLength() > 0 {
		element = (*s).stack[s.getLength()-1]
		s.stack = s.stack[:s.getLength()-1]
		return element, nil
	} else {
		return -1, errors.New("Pop() on empty stack!")
	}
}

func (s *stack) peek() (element int, err error) {
	if (*s).getLength() > 0 {
		element = (*s).stack[s.getLength()-1]
		return element, nil
	} else {
		return -1, errors.New("Peek() on empty stack!")
	}
}
