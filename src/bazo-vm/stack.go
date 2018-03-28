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

func (s *Stack) PopIndexAt(index int) (element big.Int) {
	if (*s).GetLength() >= index {
		element = (*s).stack[index]
		s.stack = append((*s).stack[:index], (*s).stack[index+1:]...)
		return element
	} else {
		log.Fatal(errors.New("Index out of bounds"))
		return *new(big.Int).SetInt64(0)
	}
}

func (s *Stack) Pop() (element big.Int) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("Pop() on empty stack"))
		return *new(big.Int).SetInt64(0)
	}
}

func (s *Stack) Peek() (element big.Int, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return *new(big.Int).SetInt64(0), errors.New("Peek() on empty stack!")
	}
}
