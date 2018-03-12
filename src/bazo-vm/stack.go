package bazo_vm

import (
	"encoding/binary"
	"errors"
	"log"
)

type byteArray []byte

type Stack struct {
	stack []byteArray
}

func NewStack() Stack {
	return Stack{
		stack: []byteArray{},
	}
}

func (s Stack) GetLength() int {
	return len(s.stack)
}

func (s *Stack) Ipush(element int) {
	ba := make(byteArray, 8)
	binary.LittleEndian.PutUint64(ba, uint64(element))
	s.stack = append(s.stack, ba)
}

func (s *Stack) Push(element byteArray) {
	s.stack = append(s.stack, element)
}

func (s *Stack) Ipop() (element int) {
	if (*s).GetLength() > 0 {
		element = int(binary.LittleEndian.Uint64((*s).stack[s.GetLength()-1]))
		s.stack = s.stack[:s.GetLength()-1]
		return element
	} else {
		log.Fatal(errors.New("Ipop() on empty stack"))
		return -1
	}
}

func (s *Stack) Ipeek() (element int, err error) {
	if (*s).GetLength() > 0 {
		element = int(binary.LittleEndian.Uint64((*s).stack[s.GetLength()-1]))
		return element, nil
	} else {
		return -1, errors.New("Ipeek() on empty stack!")
	}
}

func (s *Stack) Peek() (element byteArray, err error) {
	if (*s).GetLength() > 0 {
		element = (*s).stack[s.GetLength()-1]
		return element, nil
	} else {
		return byteArray{}, errors.New("Peek() on empty stack!")
	}
}
