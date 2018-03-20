package bazo_vm

import (
	"fmt"
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
	val := IntToByteArray(5800)
	fmt.Println(val)

	code := []byte{
		PUSH, 1, 125,
		PUSH, 2, 168, 22,
		ADD,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()
	fmt.Println(ByteArrayToInt(val))

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if reflect.DeepEqual(val, []byte{123}) {
		t.Errorf("Actual value is %v, sould be 53 after adding up 50 and 3", val)
	}
}

func TestProgramExecutionSubtraction(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 3,
		SUB,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 3 {
		t.Errorf("Actual value is %v, sould be 3 after subtracting 2 from 5", val)
	}
}

func TestProgramExecutionSubtractionWithNegativeResults(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		PUSH, 1, 6,
		SUB,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if int(ByteArrayToInt(val)) != -3 {
		t.Errorf("Actual value is %v, sould be -3 after subtracting 6 from 3", val)
	}
}

func TestProgramExecutionMultiplication(t *testing.T) {
	code := []byte{
		PUSH, 1, 5,
		PUSH, 1, 2,
		MULT,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 10 {
		t.Errorf("Actual value is %v, sould be 10 after multiplying 2 with 5", val)
	}
}

func TestProgramExecutionDivision(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 2,
		DIV,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 3 {
		t.Errorf("Actual value is %v, sould be 10 after dividing 6 by 2", val)
	}
}

func TestProgramExecutionDivisionByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic but should because divsion by 0")
		}
	}()

	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 0,
		DIV,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)
}

func TestProgramExecutionEq(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		EQ,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after comparing 4 with 4", val)
	}
}

func TestProgramExecutionNeq(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 5,
		NEQ,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after comparing 6 with 5 to not be equal", val)
	}
}

func TestProgramExecutionLt(t *testing.T) {
	code := []byte{
		PUSH, 1, 4,
		PUSH, 1, 6,
		LT,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 < 6", val)
	}
}

func TestProgramExecutionGt(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 4,
		GT,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 > 4", val)
	}
}

func TestProgramExecutionLte(t *testing.T) {
	code := []byte{
		PUSH, 1, 4,
		PUSH, 1, 6,
		LTE,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 <= 6", val)
	}

	code1 := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		LTE,
		HALT,
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 <= 6", val)
	}
}

func TestProgramExecutionGte(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 4,
		GTE,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 >= 4", val)
	}

	code1 := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		GTE,
		HALT,
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if ByteArrayToInt(val) != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 6", val)
	}
}

func TestProgramExectuionShiftl(t *testing.T) {
	code := []byte{
		PUSH, 1, 1,
		SHIFTL, 3,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	result := ByteArrayToInt(vm.evaluationStack.Pop())

	if result != 8 {
		t.Errorf("Expected result to be 8 but was %v", result)
	}
}

func TestProgramExectuionShiftr(t *testing.T) {
	code := []byte{
		PUSH, 1, 8,
		SHIFTR, 3,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	result := ByteArrayToInt(vm.evaluationStack.Pop())

	if result != 1 {
		t.Errorf("Expected result to be 1 but was %v", result)
	}
}

func TestProgramExecutionJmpif(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		PUSH, 1, 4,
		ADD,
		PUSH, 1, 20,
		LT,
		JMPIF, 16,
		PUSH, 1, 3,
		PUSHS, 0x61, 0x73, 0x64, 0x66, 0x00,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if !reflect.DeepEqual(val, []byte{0x61, 0x73, 0x64, 0x66}) {
		t.Errorf("Actual value is %v, sould be {0x61, 0x73, 0x64, 0x66} after executing program", val)
	}
}

func TestProgramExecutionJmp(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		JMP, 13,
		PUSH, 1, 4,
		ADD,
		PUSH, 1, 15,
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

	if ByteArrayToInt(val) != 3 {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", val)
	}
}

func TestProgramExecutionSha3(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		SHA3,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	val := vm.evaluationStack.Pop()

	if !reflect.DeepEqual(val, []byte{227, 237, 86, 189, 8, 109, 137, 88, 72, 58, 18, 115, 79, 160, 174, 127, 92, 139, 177, 96, 239, 144, 146, 198, 126, 130, 237, 155, 25, 228, 199, 178}) {
		t.Errorf("Actual value is %v, sould be {227, 237, 86, 189...} after jumping to halt", val)
	}
}

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
