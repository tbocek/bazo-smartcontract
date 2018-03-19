package bazo_vm

import (
	"reflect"
	"testing"
)

func TestNewVM(t *testing.T) {
	vm := NewVM(0)

	if len(vm.code) > 0 {
		t.Errorf("Actual code length is %v, should be 0 after initialization", len(vm.code))
	}

	if vm.pc != 0 {
		t.Errorf("Actual pc counter is %v, should be 0 after initialization", vm.pc)
	}
}

func TestProgramExecutionAddition(t *testing.T) {
	code := []byte{
		PUSH, 8, 125, 0, 0, 0, 0, 0, 0, 0,
		PUSH, 8, 25, 123, 0, 0, 0, 0, 0, 0,
		ADD,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if reflect.DeepEqual(val, []byte{123}) {
		t.Errorf("Actual value is %v, sould be 53 after adding up 50 and 3", val)
	}
}

/*func TestProgramExecutionSubtraction(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(5)},
		{PUSHI, IntToByteArray(2)},
		{SUB, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 3 {
		t.Errorf("Actual value is %v, sould be 3 after subtracting 2 from 5", val)
	}
}

func TestProgramExecutionMultiplication(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(5)},
		{PUSHI, IntToByteArray(2)},
		{MULT, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 10 {
		t.Errorf("Actual value is %v, sould be 10 after multiplying 2 with 5", val)
	}
}

func TestProgramExecutionDivision(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(2)},
		{DIV, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

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

	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(0)},
		{DIV, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)
}

func TestProgramExecutionEq(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(6)},
		{EQ, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after comparing 4 with 4", val)
	}
}

func TestProgramExecutionNeq(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(5)},
		{NEQ, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after comparing 6 with 5 to not be equal", val)
	}
}

func TestProgramExecutionLt(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(4)},
		{PUSHI, IntToByteArray(6)},
		{LT, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 < 6", val)
	}
}

func TestProgramExecutionGt(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(4)},
		{GT, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 > 4", val)
	}
}

func TestProgramExecutionLte(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(4)},
		{PUSHI, IntToByteArray(6)},
		{LTE, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 <= 6", val)
	}

	code1 := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(6)},
		{LTE, []byte{}},
		{HALT, []byte{}},
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, false)

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 <= 6", val)
	}
}

func TestProgramExecutionGte(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(4)},
		{GTE, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 > 4", val)
	}

	code1 := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(6)},
		{GTE, []byte{}},
		{HALT, []byte{}},
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, false)

	if val != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 6", val)
	}
}

func TestProgramExectuionShiftl(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(1)},
		{SHIFTL, IntToByteArray(3)},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	result := ByteArrayToInt(vm.evaluationStack.Pop().byteArray)

	if result != 8 {
		t.Errorf("Expected result to be 8 but was %v", result)
	}
}

func TestProgramExectuionShiftr(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(8)},
		{SHIFTR, IntToByteArray(3)},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	result := ByteArrayToInt(vm.evaluationStack.Pop().byteArray)

	if result != 1 {
		t.Errorf("Expected result to be 1 but was %v", result)
	}
}

func TestProgramExecutionJmpif(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(3)},
		{PUSHI, IntToByteArray(4)},
		{ADD, []byte{}},
		{PUSHI, IntToByteArray(15)},
		{LT, []byte{}},
		{JMPIF, IntToByteArray(7)},
		{PUSHI, IntToByteArray(456)},
		{PUSHI, IntToByteArray(10)},
		{PUSHI, IntToByteArray(10)},
		{ADD, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 20 {
		t.Errorf("Actual value is %v, sould be 20 after executing program", val)
	}
}

func TestProgramExecutionJmp(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(3)},
		{JMP, IntToByteArray(6)},
		{PUSHI, IntToByteArray(4)},
		{ADD, []byte{}},
		{PUSHI, IntToByteArray(15)},
		{ADD, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 3 {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", val)
	}
}

func TestProgramExecutionSha3(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(3)},
		{SHA3, []byte{}},
		{HALT, []byte{}},
	}

	vm := NewVM(0)
	vm.Exec(code, false)

	val := vm.evaluationStack.PopStr()

	if val != "8dfdf0627f9577b519a37dd574796c2110717e5ffe213df69c7cf0dab8853427" {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", val)
	}
}
*/
func TestProgramExecutionPushs(t *testing.T) {
	code := []byte{
		PUSHS, 0x61, 0x73, 0x64, 0x66, 0x00,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	first := vm.evaluationStack.Pop()

	if ByteArrayToString(first) != "asdf" {
		t.Errorf("Actual value is %s, sould be 'Bierchen' after popping string", first)
	}
}
