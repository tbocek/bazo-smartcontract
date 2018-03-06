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
		t.Errorf("Actual value is %v, sould be 5 after adding up 2 and 3", val)
	}
}

func TestProgramExecutionSubtraction(t *testing.T) {
	code := []int{
		PUSH, 5, //0, 5
		PUSH, 2, //0, 2
		SUB,
		HALT,
	}

	vm := NewVM()
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 3 {
		t.Errorf("Actual value is %v, sould be 3 after subtracting 2 from 5", val)
	}
}

func TestProgramExecutionMultiplication(t *testing.T) {
	code := []int{
		PUSH, 5, //0, 5
		PUSH, 2, //0, 2
		MULT,
		HALT,
	}

	vm := NewVM()
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 10 {
		t.Errorf("Actual value is %v, sould be 10 after multiplying 2 with 5", val)
	}
}

func TestProgramExecutionDivision(t *testing.T) {
	code := []int{
		PUSH, 6, //0, 6
		PUSH, 2, //0, 2
		DIV,
		HALT,
	}

	vm := NewVM()
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 3 {
		t.Errorf("Actual value is %v, sould be 10 after dividing 6 by 2", val)
	}
}

func TestProgramExecutionDivisionByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic but should because divsion by 0")
		}
	}()

	code := []int{
		PUSH, 6, //0, 6
		PUSH, 0, //0, 0
		DIV,
		HALT,
	}

	vm := NewVM()
	vm.Exec(code, true)
}
