package main

import (
	"testing"
	"encoding/binary"
)

func TestNewStack(t *testing.T) {
	s := newStack()
	if s.getSize() != 0 {
		t.Errorf("Expected stack with size 0 but got %v", s.getSize())
	}
}

func TestStackPop(t *testing.T) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 31415926)
	s := stack{
		size: 1,
		stack: []byteArray{bs},
	}

	ba, err := s.pop()
	if err != nil {
		t.Errorf("Stack threw an error")
	}

	result := binary.LittleEndian.Uint32(ba)

	if s.getSize() != 0 {
		t.Errorf("Expected steck with size 0 but got %v", s.getSize())
	}

	if result != 31415926 {
		t.Errorf("The byte array retrieved from stack is unequal to the initial byte array")
	}
}

func TestStackPopIfEmpty(t *testing.T) {
	s := newStack()
	_, err := s.pop()
	if err == nil {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestPush(t *testing.T){
	s := newStack()

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 316)

	s.push(bs)

	if s.getSize() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.getSize())
	}
}

func TestPushAndPopTogether(t *testing.T){
	s := newStack()

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 316)

	s.push(bs)
	s.push(bs)
	s.push(bs)

	if s.getSize() != 3 {
		t.Errorf("Expected size to be 1 but was %v", s.getSize())
	}

	bs2, _ := s.pop()

	if s.size != 2 {
		t.Errorf("Expected size to be 2 but was %v", s.getSize())
	}

	i := binary.LittleEndian.Uint32(bs2)

	if i != 316 {
		t.Errorf("Expected inserted Element to be 316 but was %v", i)
	}
}