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










