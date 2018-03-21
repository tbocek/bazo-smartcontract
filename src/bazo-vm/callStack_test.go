package bazo_vm

import (
	"reflect"
	"testing"
)

func TestCallStack_NewCallStack(t *testing.T) {
	cs := NewCallStack()
	if cs.GetLength() != 0 {
		t.Errorf("Expected memory with size 0 but got %v", cs.GetLength())
	}
}

func TestCallStack_Push(t *testing.T) {
	cs := NewCallStack()

	variables := map[int][]byte{
		0: {12, 39, 43},
		1: {89, 128, 12},
		2: {12, 39, 43},
	}

	cs.Push(Frame{variables: variables, returnAddress: 3})
	topOfStack := cs.Pop()

	if reflect.DeepEqual(topOfStack, variables) {
		t.Errorf("Expected same as variable defined above, but got %s", topOfStack)
	}

	if cs.GetLength() != 0 {
		t.Errorf("Expected empty stack to throw an error when using pop() but it didn't")
	}
}

func TestCallStack_MultiplePushPop(t *testing.T) {
	cs := NewCallStack()

	variables1 := map[int][]byte{
		0: {12, 39, 43},
	}

	variables2 := map[int][]byte{
		0: {12, 39, 43},
		1: {89, 128, 12},
	}

	variables3 := map[int][]byte{
		0: {12, 39, 43},
		1: {89, 128, 12},
		2: {12, 39, 43},
	}

	cs.Push(Frame{variables: variables1, returnAddress: 0})
	cs.Push(Frame{variables: variables2, returnAddress: 0})
	cs.Push(Frame{variables: variables3, returnAddress: 0})

	if cs.GetLength() != 3 {
		t.Errorf("Expected Lenght to be 3 after Pushing 3 Frames, but got %s", cs.GetLength())
	}

	cs.Pop()
	cs.Pop()

	if cs.GetLength() != 1 {
		t.Errorf("Expected Lenght to be 1 after Pushing 3 Frames and Popping twice, but got %s", cs.GetLength())
	}

	topOfStack := cs.Pop().variables
	if !reflect.DeepEqual(topOfStack, variables1) {
		t.Errorf("Expected variables popped to be %v but got %v", variables1, topOfStack)
	}
}
