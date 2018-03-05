package main

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := newStack()
	if s.getLength() != 0 {
		t.Errorf("Expected stack with size 0 but got %v", s.getLength())
	}
}

func TestStackPopWhenEmpty(t *testing.T) {
	s := newStack()
	_, err := s.pop()
	if err == nil {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestStackPopIfRemoves(t *testing.T) {
	s := newStack()

	s.push(3)
	s.pop()

	_, err := s.pop()
	if err == nil {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestStackPeek(t *testing.T) {
	s := newStack()

	s.push(3)
	s.peek()

	if s.getLength() != 1 {
		t.Errorf("Expected stack with size 1 but got %v", s.getLength())
	}
}

func TestPushAndPopElement(t *testing.T) {
	s := newStack()

	if s.getLength() != 0 {
		t.Errorf("Expected size before push to be 0, but was %v", s.getLength())
	}

	s.push(2)

	if s.getLength() != 1 {
		t.Errorf("Expected size to be 1 but was %v", s.getLength())
	}

	val, _ := s.pop()
	if val != 2 {
		t.Errorf("Expected val of element to be 2, but was %v", val)
	}

	s.push(5)

	if s.getLength() != 2 {
		t.Errorf("Expected size to be 1 but was %v", s.getLength())
	}
}
