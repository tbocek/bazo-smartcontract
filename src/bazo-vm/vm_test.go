package main

import (
	"testing"
)

func TestNewVM(t *testing.T) {
	vm := NewVM()

	if len(vm.code) > 0 {
		t.Errorf("Actual code length is %v, should be 0 after initialization", len(vm.code))
	}

	if vm.pc != 0 {
		t.Errorf("Actual pc counter is %v, should be 0 after initialization", vm.pc)
	}
}

func TestProgramExecutionAddition(t *testing.T) {
	code := []int{
		PUSH, 2, //0, 2
		PUSH, 3, //0, 3
		ADD, //1
		HALT,
	}

	vm := NewVM()
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 5 {
		t.Errorf("Actual value is %v, sould be 5 after addition", val)
	}
}
