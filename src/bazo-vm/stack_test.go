package bazo_vm

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.GetLength() != 0 {
		t.Errorf("Expected stack with size 0 but got %v", s.GetLength())
	}
}

func TestStackPopWhenEmpty(t *testing.T) {
	s := NewStack()
	val, err := s.Peek()

	if err == nil {
		t.Errorf("Throw error because val was %v", val)
	}
}

func TestStackPopIfRemoves(t *testing.T) {
	s := NewStack()

	var val1 big.Int
	val1.SetInt64(454)

	var val2 big.Int
	val2.SetInt64(46542)

	var val3 big.Int
	val3.SetInt64(-841324768)

	s.Push(val1)
	s.Push(val2)
	s.Push(val3)

	tos, _ := s.Pop()

	if tos.Int64() != int64(-841324768) {
		t.Errorf("Expected 123 got something else")
	}

	s.Pop()
	s.Pop()

	if s.GetLength() != 0 {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()

	var val big.Int
	val.SetInt64(-841324768)

	s.Push(val)
	s.Peek()

	if s.GetLength() != 1 {
		t.Errorf("Expected stack with size 1 but got %v", s.GetLength())
	}
}

func TestStack_PopIndexAt(t *testing.T) {
	s := NewStack()

	s.Push(*big.NewInt(int64(3)))
	s.Push(*big.NewInt(int64(4)))
	s.Push(*big.NewInt(int64(5)))
	s.Push(*big.NewInt(int64(6)))
	element, _ := s.PopIndexAt(2)

	fmt.Println(s)

	if s.GetLength() != 3 {
		t.Errorf("Expected stack with size 3 but got %v", s.GetLength())
	}

	if element.Int64() != 5 {
		t.Errorf("Expected element to be 5 but got %v", element)
	}
}

func TestPushAndPopElement(t *testing.T) {
	s := NewStack()

	if s.GetLength() != 0 {
		t.Errorf("Expected size before push to be 0, but was %v", s.GetLength())
	}

	s.Push(*big.NewInt(int64(2)))

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	val, _ := s.Pop()
	if val.Int64() != 2 {
		t.Errorf("Expected val of element to be 2, but was %v", val)
	}

	s.Push(*big.NewInt(int64(5)))

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	fmt.Print(s)
}
