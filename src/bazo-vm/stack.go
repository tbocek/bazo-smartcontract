package bazo_vm

import (
	"errors"
	"math/big"
)

type Stack struct {
	stack       []big.Int
	memoryUsage uint32 // In bytes
	memoryMax   uint32
}

func NewStack() *Stack {
	return &Stack{
		stack:       nil,
		memoryUsage: 0,
		memoryMax:   1000000, // Max 1000000 Bytes = 1MB
	}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(element big.Int) {
	s.memoryUsage += getElementMemoryUsage(element.BitLen())
	s.stack = append(s.stack, element)
}

func (s *Stack) PopIndexAt(index int) (element big.Int, err error) {
	if (*s).GetLength() >= index {
		element = (*s).stack[index]
		s.memoryUsage -= getElementMemoryUsage(element.BitLen())
		s.stack = append((*s).stack[:index], (*s).stack[index+1:]...)
		return element, nil
	} else {
		return *new(big.Int).SetInt64(0), errors.New("index out of bounds")
	}
}

func (s *Stack) Pop() (element big.Int, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.memoryUsage -= getElementMemoryUsage(element.BitLen())
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

// Function turns bit into bytes and rounds up
func getElementMemoryUsage(element int) uint32 {
	return uint32(((element + 7) / 8) + 1)
}
