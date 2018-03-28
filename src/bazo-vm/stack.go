package bazo_vm

import (
	"errors"
	"log"
	"math/big"
)

type Stack struct {
	stack []big.Int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(element big.Int) {
	s.stack = append(s.stack, element)
}

func (s *Stack) PopIndexAt(index int) (element big.Int, err error) {
	if (*s).GetLength() >= index {
		element = (*s).stack[index]
		s.stack = append((*s).stack[:index], (*s).stack[index+1:]...)
		return element, nil
	} else {
		return *new(big.Int).SetInt64(0), errors.New("index out of bounds")
	}
}

func (s *Stack) Pop() (element big.Int, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element, nil
	} else {
		return *new(big.Int).SetInt64(0), errors.New("pop() on empty stack")
	}
}

func (s *Stack) Peek() (element big.Int, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return *new(big.Int).SetInt64(0), errors.New("peek() on empty stack")
	}
}
