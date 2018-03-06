package main

import (
	"errors"
	"log"
)

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{
		stack: []int{},
	}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(element int) {
	s.stack = append(s.stack, element)
}

func (s *Stack) Pop() (element int) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("Pop() on empty stack"))
		return -1
	}
}

func (s *Stack) Peek() (element int, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return -1, errors.New("Peek() on empty stack!")
	}
}
