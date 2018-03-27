package bazo_vm

import (
	"fmt"
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

	int := IntToByteArray(123)
	s.Push(int)
	asdf := ByteArrayToInt(s.Pop())

	if asdf != 123 {
		t.Errorf("Expected 123 got something else")

	}

	if s.GetLength() != 0 {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()

	s.Push(IntToByteArray(3))
	s.Peek()

	if s.GetLength() != 1 {
		t.Errorf("Expected stack with size 1 but got %v", s.GetLength())
	}
}

func TestStack_PopIndexAt(t *testing.T) {
	s := NewStack()

	s.Push(IntToByteArray(3))
	s.Push(IntToByteArray(4))
	s.Push(IntToByteArray(5))
	s.Push(IntToByteArray(6))
	s.PopIndexAt(2)
	s.Peek()

	fmt.Println(s)

	if s.GetLength() != 3 {
		t.Errorf("Expected stack with size 3 but got %v", s.GetLength())
	}
}

func TestPushAndPopElement(t *testing.T) {
	s := NewStack()

	if s.GetLength() != 0 {
		t.Errorf("Expected size before push to be 0, but was %v", s.GetLength())
	}

	s.Push(IntToByteArray(2))

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	val := ByteArrayToInt(s.Pop())
	if val != 2 {
		t.Errorf("Expected val of element to be 2, but was %v", val)
	}

	s.Push(IntToByteArray(5))

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	fmt.Print(s)
}
