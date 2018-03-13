package bazo_vm

import (
	"fmt"
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
	code := []instruction{
		{PUSHI, IntToByteArray(50)},
		{PUSHI, IntToByteArray(3)},
		{ADD, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	// Get evaluationStack top value to compare to expected value
	val, err := vm.evaluationStack.PeekInt()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if val != 53 {
		t.Errorf("Actual value is %v, sould be 53 after adding up 50 and 3", val)
	}
	fmt.Println(val)
}

func TestProgramExecutionSubtraction(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(5)},
		{PUSHI, IntToByteArray(2)},
		{SUB, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{MULT, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{DIV, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{DIV, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)
}

func TestProgramExecutionEq(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(6)},
		{EQ, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{NEQ, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{LT, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{GT, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{LTE, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{LTE, byteArray{}},
		{HALT, byteArray{}},
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 <= 6", val)
	}
}

func TestProgramExecutionGte(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(6)},
		{PUSHI, IntToByteArray(4)},
		{GTE, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{GTE, byteArray{}},
		{HALT, byteArray{}},
	}

	vm1 := NewVM(0)
	vm1.Exec(code1, true)

	if val != 1 {
		t.Errorf("Actual value is %v, sould be 1 after evaluating 6 >= 6", val)
	}
}

func TestProgramExecutionJmpif(t *testing.T) {
	code := []instruction{
		{PUSHI, IntToByteArray(3)},
		{PUSHI, IntToByteArray(4)},
		{ADD, byteArray{}},
		{PUSHI, IntToByteArray(15)},
		{LT, byteArray{}},
		{JMPIF, IntToByteArray(7)},
		{PUSHI, IntToByteArray(456)},
		{PUSHI, IntToByteArray(10)},
		{PUSHI, IntToByteArray(10)},
		{ADD, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{ADD, byteArray{}},
		{PUSHI, IntToByteArray(15)},
		{ADD, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

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
		{SHA3, byteArray{}},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	val := vm.evaluationStack.PopStr()
	fmt.Println(val)

	if val != "8dfdf0627f9577b519a37dd574796c2110717e5ffe213df69c7cf0dab8853427" {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", val)
	}
}

func TestProgramExecutionPushs(t *testing.T) {
	code := []instruction{
		{PUSHS, StrToByteArray("Lecker")},
		{PUSHS, StrToByteArray("Bierchen")},
		{PUSHS, StrToByteArray("trinken")},
		{HALT, byteArray{}},
	}

	vm := NewVM(0)
	vm.Exec(code, true)

	vm.evaluationStack.PopStr()
	second := vm.evaluationStack.PopStr()

	if second != "Bierchen" {
		t.Errorf("Actual value is %s, sould be 'Bierchen' after popping string", second)
	}
}
