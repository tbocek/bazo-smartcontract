package bazo_vm

import (
	"reflect"
	"testing"
)

func newTestContextObj() Context {
	data := map[int][]byte{}

	return Context{
		transactionSender:    []byte{},
		transactioninputData: []byte{},
		maxGasAmount:         100000,
		smartContract:        NewSmartContract([]byte{}, 100, true, []byte{}, []byte{}, data),
	}
}

func TestVMGasConsumption(t *testing.T) {
	vm := NewVM(0)

	context := newTestContextObj()
	context.maxGasAmount = 3

	code := []byte{
		PUSH, 1, 8,
		PUSH, 1, 8,
		ADD,
		HALT,
	}

	context.smartContract.data.code = code

	vm.Exec(context, true)
	ba, _ := vm.evaluationStack.Pop()
	val := ba

	if val.Int64() != 16 {
		t.Errorf("Expected first value to be 16 but was %v", val)
	}
}

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
		PUSH, 1, 125,
		PUSH, 2, 168, 22,
		ADD,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != int64(43155) {
		t.Errorf("Actual value is %v, should be 53 after adding up 50 and 3", tos.Int64())
	}
}

func TestProgramExecutionSubtraction(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 3,
		SUB,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 3 after subtracting 2 from 5", tos)
	}
}

func TestProgramExecutionSubtractionWithNegativeResults(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		PUSH, 1, 6,
		SUB,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != -3 {
		t.Errorf("Actual value is %v, should be -3 after subtracting 6 from 3", tos)
	}
}

func TestProgramExecutionMultiplication(t *testing.T) {
	code := []byte{
		PUSH, 1, 5,
		PUSH, 1, 2,
		MULT,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 10 {
		t.Errorf("Actual value is %v, should be 10 after multiplying 2 with 5", tos)
	}
}

func TestProgramExecutionDivision(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 2,
		DIV,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 10 after dividing 6 by 2", tos)
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

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)
}

func TestProgramExecutionEq(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		EQ,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after comparing 4 with 4", tos)
	}
}

func TestProgramExecutionNeq(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 5,
		NEQ,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after comparing 6 with 5 to not be equal", tos)
	}
}

func TestProgramExecutionLt(t *testing.T) {
	code := []byte{
		PUSH, 1, 4,
		PUSH, 1, 6,
		LT,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 4 < 6", tos)
	}
}

func TestProgramExecutionGt(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 4,
		GT,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 > 4", tos)
	}
}

func TestProgramExecutionLte(t *testing.T) {
	code := []byte{
		PUSH, 1, 4,
		PUSH, 1, 6,
		LTE,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 4 <= 6", tos)
	}

	code1 := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		LTE,
		HALT,
	}

	vm1 := NewVM(0)
	context1 := newTestContextObj()
	context1.smartContract.data.code = code1
	vm1.Exec(context1, true)

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 <= 6", tos)
	}
}

func TestProgramExecutionGte(t *testing.T) {
	code := []byte{
		PUSH, 1, 6,
		PUSH, 1, 4,
		GTE,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 4", tos)
	}

	code1 := []byte{
		PUSH, 1, 6,
		PUSH, 1, 6,
		GTE,
		HALT,
	}

	context1 := newTestContextObj()
	context1.smartContract.data.code = code1

	vm1 := NewVM(0)
	vm1.Exec(context1, true)

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 6", tos)
	}
}

func TestProgramExectuionShiftl(t *testing.T) {
	code := []byte{
		PUSH, 1, 1,
		SHIFTL, 3,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 8 {
		t.Errorf("Expected result to be 8 but was %v", tos)
	}
}

func TestProgramExectuionShiftr(t *testing.T) {
	code := []byte{
		PUSH, 1, 8,
		SHIFTR, 3,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 1 {
		t.Errorf("Expected result to be 1 but was %v", tos)
	}
}

func TestProgramExecutionJmpif(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		PUSH, 1, 4,
		ADD,
		PUSH, 1, 20,
		LT,
		JMPIF, 17,
		PUSH, 1, 3,
		NOP,
		NOP,
		NOP,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	if vm.evaluationStack.GetLength() != 0 {
		t.Errorf("After calling and returning, callStack lenght should be 0, but is %v", vm.evaluationStack.GetLength())
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

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 3 after jumping to halt", tos)
	}
}

func TestProgramExecutionCall(t *testing.T) {
	code := []byte{
		PUSH, 1, 10,
		PUSH, 1, 8,
		CALL, 13, 2,
		HALT,
		NOP,
		NOP,
		LOAD, 0, // Begin of called function at address 13
		LOAD, 1,
		SUB,
		PRINT,
		RET,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if tos.Int64() != 2 {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", tos)
	}

	callStackLenght := vm.callStack.GetLength()

	if callStackLenght != 0 {
		t.Errorf("After calling and returning, callStack lenght should be 0, but is %v", callStackLenght)
	}
}

func TestProgramExecutionCallExt(t *testing.T) {
	code := []byte{
		PUSH, 1, 10,
		PUSH, 1, 8,
		CALLEXT, 227, 237, 86, 189, 8, 109, 137, 88, 72, 58, 18, 115, 79, 160, 174, 127, 92, 139, 177, 96, 239, 144, 146, 198, 126, 130, 237, 155, 25, 228, 199, 178, 41, 24, 45, 14, 2,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)
}

func TestProgramExecutionSha3(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		SHA3,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	val, _ := vm.evaluationStack.Pop()

	if !reflect.DeepEqual(val.Bytes(), []byte{227, 237, 86, 189, 8, 109, 137, 88, 72, 58, 18, 115, 79, 160, 174, 127, 92, 139, 177, 96, 239, 144, 146, 198, 126, 130, 237, 155, 25, 228, 199, 178}) {
		t.Errorf("Actual value is %v, should be {227, 237, 86, 189...} after jumping to halt", val)
	}
}

func TestProgramExecutionRoll(t *testing.T) {
	code := []byte{
		PUSH, 1, 3,
		PUSH, 1, 4,
		PUSH, 1, 5,
		PUSH, 1, 6,
		PUSH, 1, 7,
		ROLL, 2,
		HALT,
	}

	context := newTestContextObj()
	context.smartContract.data.code = code

	vm := NewVM(0)
	vm.Exec(context, true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 4 {
		t.Errorf("Actual value is %v, should be 4 after rolling with two as arg", tos)
	}
}
