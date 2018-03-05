package main

import "errors"

type byteArray []byte

type stack struct {
	size int
	//TODO maybe allocate an array with a certain size
	stack []byteArray
}

func (s stack) getSize() int{
	return s.size
}

func newStack() stack{
	return stack{
		size: 0,
		stack: []byteArray{},
	}
}

func (s *stack) push(element byteArray){
	//TODO Probably better to use a Linked list
	s.stack = append(s.stack, element)
	s.size = s.size + 1
}

func (s *stack) pop() (ba byteArray, err error) {
	if (*s).size > 0 {
		(*s).size = (*s).size - 1
		ba = (*s).stack[(*s).size]
		return ba, err
	} else {
		return []byte{}, errors.New("Pop() on empty stack!")
	}

}