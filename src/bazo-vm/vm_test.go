package bazo_vm

import (
	"reflect"
	"testing"
	"math/big"
)

func TestVMGasConsumption(t *testing.T) {
	code := []byte{
		PUSH, 0, 8,
		PUSH, 0, 8,
		ADD,
		HALT,
	}

	vm := NewVM()
	vm.context.maxGasAmount = 3
	vm.context.contractAccount.Code = code

	vm.Exec(true)
	ba, _ := vm.evaluationStack.Pop()
	val := ba

	if val.Int64() != 16 {
		t.Errorf("Expected first value to be 16 but was %v", val)
	}
}

func TestNewVM(t *testing.T) {
	vm := NewVM()

	if len(vm.code) > 0 {
		t.Errorf("Actual code length is %v, should be 0 after initialization", len(vm.code))
	}

	if vm.pc != 0 {
		t.Errorf("Actual pc counter is %v, should be 0 after initialization", vm.pc)
	}
}

func TestAddition(t *testing.T) {
	code := []byte{
		PUSH, 0, 125,
		PUSH, 1, 168, 22,
		ADD,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != int64(43155) {
		t.Errorf("Actual value is %v, should be 53 after adding up 50 and 3", tos.Int64())
	}
}

func TestSubtraction(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 3,
		SUB,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 3 after subtracting 2 from 5", tos)
	}
}

func TestSubtractionWithNegativeResults(t *testing.T) {
	code := []byte{
		PUSH, 0, 3,
		PUSH, 0, 6,
		SUB,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != -3 {
		t.Errorf("Actual value is %v, should be -3 after subtracting 6 from 3", tos)
	}
}

func TestMultiplication(t *testing.T) {
	code := []byte{
		PUSH, 0, 5,
		PUSH, 0, 2,
		MULT,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 10 {
		t.Errorf("Actual value is %v, should be 10 after multiplying 2 with 5", tos)
	}
}

func TestModulo(t *testing.T) {
	code := []byte{
		PUSH, 0, 5,
		PUSH, 0, 2,
		MOD,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after 5 mod 2", tos)
	}
}

func TestNegate(t *testing.T) {
	code := []byte{
		PUSH, 0, 5,
		NEG,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != -5 {
		t.Errorf("Actual value is %v, should be -5 after negating 5", tos)
	}
}

func TestDivision(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 2,
		DIV,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 10 after dividing 6 by 2", tos)
	}
}

func TestDivisionByZero(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 0,
		DIV,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	result, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("%v", err)
	}

	e := BigIntToString(result)
	if e != "Division by Zero" {
		t.Errorf("Expected Error Message to be returned but got: %v", e)
	}
}

func TestEq(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 6,
		EQ,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after comparing 4 with 4", tos)
	}
}

func TestNeq(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 5,
		NEQ,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after comparing 6 with 5 to not be equal", tos)
	}
}

func TestLt(t *testing.T) {
	code := []byte{
		PUSH, 0, 4,
		PUSH, 0, 6,
		LT,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 4 < 6", tos)
	}
}

func TestGt(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 1, 0, 4,
		GT,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 > 4", tos)
	}
}

func TestLte(t *testing.T) {
	code := []byte{
		PUSH, 0, 4,
		PUSH, 0, 6,
		LTE,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 4 <= 6", tos)
	}

	code1 := []byte{
		PUSH, 0, 6,
		PUSH, 0, 6,
		LTE,
		HALT,
	}

	vm1 := NewVM()
	vm1.context.contractAccount.Code = code1
	vm1.Exec(true)

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 <= 6", tos)
	}
}

func TestGte(t *testing.T) {
	code := []byte{
		PUSH, 0, 6,
		PUSH, 0, 4,
		GTE,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 4", tos)
	}

	code1 := []byte{
		PUSH, 0, 6,
		PUSH, 0, 6,
		GTE,
		HALT,
	}

	vm1 := NewVM()
	vm1.context.contractAccount.Code = code1
	vm1.Exec(true)

	if tos.Int64() != 1 {
		t.Errorf("Actual value is %v, should be 1 after evaluating 6 >= 6", tos)
	}
}

func TestShiftl(t *testing.T) {
	code := []byte{
		PUSH, 0, 1,
		SHIFTL, 3,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 8 {
		t.Errorf("Expected result to be 8 but was %v", tos)
	}
}

func TestShiftr(t *testing.T) {
	code := []byte{
		PUSH, 0, 8,
		SHIFTR, 3,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 1 {
		t.Errorf("Expected result to be 1 but was %v", tos)
	}
}

func TestJmpif(t *testing.T) {
	code := []byte{
		PUSH, 0, 3,
		PUSH, 0, 4,
		ADD,
		PUSH, 0, 20,
		LT,
		JMPIF, 17,
		PUSH, 0, 3,
		NOP,
		NOP,
		NOP,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	if vm.evaluationStack.GetLength() != 0 {
		t.Errorf("After calling and returning, callStack lenght should be 0, but is %v", vm.evaluationStack.GetLength())
	}
}

func TestJmp(t *testing.T) {
	code := []byte{
		PUSH, 0, 3,
		JMP, 13,
		PUSH, 0, 4,
		ADD,
		PUSH, 0, 15,
		ADD,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	if tos.Int64() != 3 {
		t.Errorf("Actual value is %v, should be 3 after jumping to halt", tos)
	}
}

func TestCall(t *testing.T) {
	code := []byte{
		PUSH, 0, 10,
		PUSH, 0, 8,
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

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

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

func TestCallExt(t *testing.T) {
	code := []byte{
		PUSH, 0, 10,
		PUSH, 0, 8,
		CALLEXT, 227, 237, 86, 189, 8, 109, 137, 88, 72, 58, 18, 115, 79, 160, 174, 127, 92, 139, 177, 96, 239, 144, 146, 198, 126, 130, 237, 155, 25, 228, 199, 178, 41, 24, 45, 14, 2,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)


}

func TestSload(t *testing.T){
	code := []byte{
		SLOAD, 0,
		HALT,

	}

	vm := NewVM()
	vm.context.contractAccount.Code = code

	//TODO Contract Variables should not be modifyable in the VM only after execution
	variable := []big.Int{}
	vm.context.contractAccount.ContractVariables = append(variable, StrToBigInt("Hi There!!"))
	vm.Exec(true)

	result, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("%v", err)
	}

	resultString := BigIntToString(result)
	if resultString != "Hi There!!" {
		t.Errorf("The String on the Stack should be 'Hi There!!' but was %v", resultString)
	}
}

func TestSstore(t *testing.T){
	code := []byte{
		PUSH, 9, 72, 105, 32, 84, 104, 101, 114, 101, 33, 33,
		SSTORE, 0,
		HALT,

	}

	vm := NewVM()
	vm.context.contractAccount.Code = code

	//TODO Contract Variables should not be modifyable in the VM only after execution
	variable := []big.Int{StrToBigInt("Something")}
	vm.context.contractAccount.ContractVariables = variable
	vm.Exec(true)

	result := BigIntToString(vm.context.contractAccount.ContractVariables[0])
	if result != "Hi There!!" {
		t.Errorf("The String on the Stack should be 'Hi There!!' but was '%v'", result)
	}
}

func TestSha3(t *testing.T) {
	code := []byte{
		PUSH, 0, 3,
		SHA3,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	val, _ := vm.evaluationStack.Pop()

	if !reflect.DeepEqual(val.Bytes(), []byte{227, 237, 86, 189, 8, 109, 137, 88, 72, 58, 18, 115, 79, 160, 174, 127, 92, 139, 177, 96, 239, 144, 146, 198, 126, 130, 237, 155, 25, 228, 199, 178}) {
		t.Errorf("Actual value is %v, should be {227, 237, 86, 189...} after jumping to halt", val)
	}
}

func TestRoll(t *testing.T) {
	code := []byte{
		PUSH, 0, 3,
		PUSH, 0, 4,
		PUSH, 0, 5,
		PUSH, 0, 6,
		PUSH, 0, 7,
		ROLL, 2,
		HALT,
	}

	vm := NewVM()
	vm.context.contractAccount.Code = code
	vm.Exec(true)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 4 {
		t.Errorf("Actual value is %v, should be 4 after rolling with two as arg", tos)
	}
}
