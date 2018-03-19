package bazo_vm

import (
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

/*
func TestStackPopStr(t *testing.T) {
	s := NewStack()

	s.PushStr("asdfasdf")
	val := s.PopStr()

	if val != "asdfasdf" {
		t.Errorf("Expected asdfasdf but value was %v ", val)
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()

	s.PushInt(3)
	s.PeekInt()

	if s.GetLength() != 1 {
		t.Errorf("Expected stack with size 1 but got %v", s.GetLength())
	}
}

func TestPushAndPopElement(t *testing.T) {
	s := NewStack()

	if s.GetLength() != 0 {
		t.Errorf("Expected size before push to be 0, but was %v", s.GetLength())
	}

	s.PushInt(2)

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	val := s.PopInt()
	if val != 2 {
		t.Errorf("Expected val of element to be 2, but was %v", val)
	}

	s.PushInt(5)

	if s.GetLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.GetLength())
	}

	fmt.Print(s)
}

*/
