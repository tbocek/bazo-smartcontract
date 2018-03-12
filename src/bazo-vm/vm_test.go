package bazo_vm

import (
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
	code := []int{
		PUSH, 2, //0, 2
		PUSH, 3, //0, 3
		ADD, //1
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

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

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

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

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

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

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

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

	vm := NewVM(0)
	vm.Exec(code, true)
}

func TestProgramExecutionEq(t *testing.T) {
	code := []int{
		PUSH, 4, //0, 4
		EQ, 4,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after comparing 4 with 4", val)
	}
}

func TestProgramExecutionNeq(t *testing.T) {
	code := []int{
		PUSH, 4, //0, 4
		NEQ, 4,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 0 {
		t.Errorf("Actual value is %v, sould be 0 after comparing 4 with 4", val)
	}
}

func TestProgramExecutionLt(t *testing.T) {
	code := []int{
		PUSH, 4, //0, 4
		LT, 6,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 < 6", val)
	}
}

func TestProgramExecutionGt(t *testing.T) {
	code := []int{
		PUSH, 6, //0, 4
		GT, 4,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 > 4", val)
	}
}

func TestProgramExecutionLte(t *testing.T) {
	code := []int{
		PUSH, 4, //0, 4
		LTE, 6,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 4 <= 6", val)
	}

	code1 := []int{
		PUSH, 6, //0, 4
		LTE, 6,
		HALT,
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 <= 6", val)
	}
}

func TestProgramExecutionGte(t *testing.T) {
	code := []int{
		PUSH, 6, //0, 4
		GTE, 4,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 > 4", val)
	}

	code1 := []int{
		PUSH, 6, //0, 4
		GTE, 6,
		HALT,
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 >= 6", val)
	}
}

func TestProgramExecutionJmpif(t *testing.T) {
	code := []int{
		PUSH, 3,
		PUSH, 4,
		ADD,
		LT, 15,
		JMPIF, 2,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 15 {
		t.Errorf("Actual value is %v, sould be 23 after executing program", val)
	}
}

func TestProgramExecutionJmp(t *testing.T) {
	code := []int{
		PUSH, 3,
		JMP, 10,
		PUSH, 4,
		ADD,
		PUSH, 15,
		ADD,
		HALT,
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.Ipeek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 3 {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", val)
	}
}