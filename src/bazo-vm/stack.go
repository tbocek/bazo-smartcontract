package bazo_vm

import (
	"errors"
)

type byteArray []byte

type Stack struct {
	stack []byteArray
}

func NewStack() *Stack {
	return &Stack{}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Push(element []byte) {
	s.stack = append(s.stack, element)
}

func (s *Stack) PopIndexAt(index int) (element []byte, err error) {
	if (*s).GetLength() >= index {
		element = (*s).stack[index]
		s.stack = append((*s).stack[:index], (*s).stack[index+1:]...)
		return element, nil
	} else {
		return []byte{}, errors.New("index out of bounds")
	}
}

func (s *Stack) Pop() (element []byte, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		s.stack = s.stack[:s.GetLength()-1]
		return element, nil
	} else {
		return []byte{}, errors.New("pop() on empty stack")
	}
}

func (s *Stack) Peek() (element []byte, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return []byte{}, errors.New("peek() on empty stack")
	}
}
