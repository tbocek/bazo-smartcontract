package main

import (
	"testing"
)

func TestNewVM(t *testing.T) {
	program := []byte{0x00, 0x02, 0x04}
	vm := newVM(program)

	if vm.getProgram()[0] != 0x00 {
		t.Errorf("Expected first program instruction to be 0x00 but was %v", vm.getProgram()[0])
	}
}

func TestProgramExecutionAddition(t *testing.T) {
	program := []byte{PUSH1, PUSH1, ADD, POP} // 1, 1, +, Pop
	vm := newVM(program)

	result := vm.run()

	if result != 0 {
		t.Errorf("Expected result to be 0 but was %v", result)
	}
}

func TestProgramExecutionSubtraction(t *testing.T) {
	program := []byte{PUSH1, PUSH1, SUB, POP} // 1, 1, -, Pop
	vm := newVM(program)

	result := vm.run()

	if result != 0 {
		t.Errorf("Expected result to be 0 but was %v", result)
	}
}
