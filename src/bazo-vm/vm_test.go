package bazo_vm

import (
	"math/big"
	"reflect"
	"testing"
	"bytes"
)

func TestVMGasConsumption(t *testing.T) {
	code := []byte{
		PUSH, 0, 8,
		PUSH, 0, 8,
		ADD,
		HALT,
	}

	vm := NewVM()
	vm.context.MaxGasAmount = 3
	vm.context.ContractAccount.Contract = code

	vm.Exec(false)
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

func TestPushOutOfBounds(t *testing.T) {
	code := []byte{
		PUSH, 0, 125,
		PUSH, 126, 12,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("%v", err)
	}

	e := BigIntToString(tos)
	if e != "instructionSet out of bounds" {
		t.Errorf("Expected Error Message to be returned but got: %v", e)
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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm1.context.ContractAccount.Contract = code1
	vm1.context.MaxGasAmount = 50
	vm1.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm1.context.ContractAccount.Contract = code1
	vm1.context.MaxGasAmount = 50
	vm1.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
		RET,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

	tos, err := vm.evaluationStack.Peek()

	if err != nil {
		t.Errorf("Expected empty Stack to throw an error when using peek() but it didn't")
	}

	if tos.Int64() != 2 {
		t.Errorf("Actual value is %v, sould be 3 after jumping to halt", tos)
	}

	callStackLenght := vm.callStack.GetLength()

	if callStackLenght != 0 {
		t.Errorf("After calling and returning, callStack lenght should be 0, but is %v", callStackLenght)
	}
}

func TestTosSize(t *testing.T) {
	code := []byte{
		PUSH, 2, 10, 4, 5,
		SIZE,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

	tos, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("Expected empty stack to throw an error when using peek() but it didn't")
	}

	if tos.Int64() != 4 {
		t.Errorf("Expected TOS size to be 4, but got %v", tos)
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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

}

func TestSload(t *testing.T) {
	code := []byte{
		SLOAD, 0,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code

	//TODO Contract Variables should not be modifyable in the VM only after execution
	variable := []big.Int{}
	vm.context.ContractAccount.ContractVariables = append(variable, StrToBigInt("Hi There!!"))
	vm.Exec(false)

	result, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("%v", err)
	}

	resultString := BigIntToString(result)
	if resultString != "Hi There!!" {
		t.Errorf("The String on the Stack should be 'Hi There!!' but was %v", resultString)
	}
}

func TestSstore(t *testing.T) {
	code := []byte{
		PUSH, 9, 72, 105, 32, 84, 104, 101, 114, 101, 33, 33,
		SSTORE, 0,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code

	//TODO Contract Variables should not be modifyable in the VM only after execution
	variable := []big.Int{StrToBigInt("Something")}
	vm.context.ContractAccount.ContractVariables = variable
	vm.Exec(false)

	result := BigIntToString(vm.context.ContractAccount.ContractVariables[0])
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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

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
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if tos.Int64() != 4 {
		t.Errorf("Actual value is %v, should be 4 after rolling with two as arg", tos)
	}
}

func TestNewMap(t *testing.T){
	code := []byte{
		NEWMAP,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.Exec(true)

	r, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("%v", err)
	}
	result := r.Bytes()
	expected := []byte{0x01, 0x00, 0x00,}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected the Value of the new Map to be %v but was %v", expected, result)
	}
}

func TestMapPush(t *testing.T){
	code := []byte{
		NEWMAP,
		PUSH, 1, 72, 105,
		PUSH, 0, 0x03,
		MAPPUSH,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	m, err := vm.evaluationStack.Pop()
	if err != nil {
		t.Errorf("%v", err)
	}

	mp, err2 := MapFromBigInt(m)
	if err2 != nil {
		t.Errorf("%v", err)
	}

	datastructure := mp[0]
	size := mp.getSize()

	if datastructure != 0x01 {
		t.Errorf("Invalid Datastructure ID, Expected 0x01 but was %v", datastructure)
	}

	if size != 1 {
		t.Errorf("invalid size, Expected 1 but was %v", size)
	}


}

func TestMapGetVAL(t *testing.T){
	code := []byte{
		NEWMAP,
		PUSH, 0x01, 0x48, 0x69,
		PUSH, 0x00, 0x03,
		MAPPUSH,
		PUSH, 0x01, 0x69, 0x69,
		PUSH, 0x00, 0x02,
		MAPPUSH,
		PUSH, 0x01, 0x48, 0x48,
		PUSH, 0x00, 0x01,
		MAPPUSH,
		PUSH, 0x00, 0x01,
		MAPGETVAL,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	v, err := vm.evaluationStack.Pop()

	if err != nil {
		t.Errorf("%v", err)
	}

	e := []byte{72, 72}
	if bytes.Compare(v.Bytes(), e) != 0 {
		t.Errorf("invalid value, Expected %v but was '%v'", e, v)
	}
}


func TestMapRemove(t *testing.T){
	code := []byte{
		NEWMAP,
		PUSH, 0x01, 0x48, 0x69,
		PUSH, 0x00, 0x03,
		MAPPUSH,
		PUSH, 0x01, 0x69, 0x69,
		PUSH, 0x00, 0x02,
		MAPPUSH,
		PUSH, 0x01, 0x48, 0x48,
		PUSH, 0x00, 0x01,
		MAPPUSH,
		PUSH, 0x00, 0x03,
		MAPREMOVE,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	mbi, err := vm.evaluationStack.Pop()
	if err != nil {
		t.Errorf("%v", err)
	}
	m, err := MapFromBigInt(mbi)
	if err != nil {
		t.Errorf("%v", err)
	}

	e := []byte{0x01,
				0x02, 0x00,
				0x01, 0x00,		0x02,
				0x02, 0x00,		0x69, 0x69,
				0x01, 0x00,		0x01,
				0x02, 0x00,		0x48, 0x48,
	}


	if bytes.Compare(m, e) != 0 {
		t.Errorf("invalid datastructure, Expected %v but was '%v'", e, m)
	}
}

func TestNewArr(t *testing.T){
	code := []byte{
		NEWARR,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	arr, err := vm.evaluationStack.Pop()
	ba := arr.Bytes()
	if err != nil {
		t.Errorf("%v", err)
	}
	expectedSize := []byte{0x00, 0x00,}
	actualSize := ba[1:3]
	if !reflect.DeepEqual(expectedSize, actualSize) {
		t.Errorf("invalid size, Expected %v but was '%v'", expectedSize, actualSize)
	}
}

func TestArrAppend(t *testing.T) {
	code := []byte{
		NEWARR,
		PUSH, 0x01, 0xFF, 0x00,
		ARRAPPEND,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)
	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	arr, err := vm.evaluationStack.Pop()
	if err != nil {
		t.Errorf("%v", err)
	}

	actual := arr.Bytes()[5:7]
	expected := []byte{0xFF, 0x00,}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("invalid element appended, Expected '%v' but was '%v'", expected, actual)
	}
}

func TestArrRemove(t *testing.T) {
	code := []byte{
		NEWARR,
		PUSH, 0x01, 0xFF, 0x00,
		ARRAPPEND,
		PUSH, 0x01, 0xAA, 0x00,
		ARRAPPEND,
		PUSH, 0x01, 0xBB, 0x00,
		ARRAPPEND,
		ARRREMOVE, 0x01, 0x00,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	a, err := vm.evaluationStack.Pop()
	if err != nil {
		t.Errorf("%v", err)
	}
	arr, bierr := ArrayFromBigInt(a)
	if bierr != nil {
		t.Errorf("%v", err)
	}

	size := arr.getSize()
	if size != uint16(2) {
		t.Errorf("invalid array size, Expected 2 but was '%v'", size)
	}

	expectedSecondElement := []byte{0xBB, 0x00,}
	actualSecondElement, err2 := arr.At(uint16(1))
	if err2 != nil {
		t.Errorf("%v", err)
	}

	if !reflect.DeepEqual(expectedSecondElement, actualSecondElement){
		t.Errorf("invalid element on second index, Expected %# x but was %# x", expectedSecondElement, actualSecondElement)
	}
}

func TestArrAt(t *testing.T) {
	code := []byte{
		NEWARR,
		PUSH, 0x01, 0xFF, 0x00,
		ARRAPPEND,
		PUSH, 0x01, 0xAA, 0x00,
		ARRAPPEND,
		PUSH, 0x01, 0xBB, 0x00,
		ARRAPPEND,
		ARRAT, 0x02, 0x00,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	exec := vm.Exec(true)

	if !exec {
		errorMessage, _ := vm.evaluationStack.Pop()
		t.Errorf("VM.Exec terminated with Error: %v", BigIntToString(errorMessage))
	}

	v1, err1 := vm.evaluationStack.Pop()

	if err1 != nil {
		t.Errorf("%v", err1)
	}

	expected1 := []byte{0xBB, 0x00}
	actual1 := v1.Bytes()
	if !reflect.DeepEqual(expected1, actual1){
		t.Errorf("invalid element on first index, Expected %# x but was %# x", expected1, actual1)
	}

}

func TestNonValidOpCode(t *testing.T) {

	code := []byte{
		89,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "Not a valid opCode" {
		t.Errorf("Expected tos to be 'Not a valid opCode' error message but was %v", tos)
	}
}

func TestArgumentsExceedInstructionSet(t *testing.T) {

	code := []byte{
		PUSH, 0x00, 0x00, PUSH, 0x0b, 0x01, 0x00, 0x03, 0x12, 0x05,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "instructionSet out of bounds" {
		t.Errorf("Expected tos to be 'instructionSet out of bounds' error message but was %v", tos)
	}
}


/*func TestPopOnEmptyStack(t *testing.T) {

	code := []byte{
		PUSH, 0x00, 0x01, SHA3, 0x05, 0x02, 0x03,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "pop() on empty stack" {
		t.Errorf("Expected tos to be 'pop() on empty stack' error message but was %v", tos)
	}
}*/

func TestFuzzReproductionInstructionSetOutOfBounds(t *testing.T) {

	code := []byte{
		PUSH, 0, 20,
		ROLL, 0,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "instructionSet out of bounds" {
		t.Errorf("Expected tos to be 'pop() on empty stack' error message but was %v", tos)
	}
}

func TestFuzzReproductionInstructionSetOutOfBounds2(t *testing.T) {

	code := []byte{
		CALLEXT, 231,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "instructionSet out of bounds" {
		t.Errorf("Expected tos to be 'pop() on empty stack' error message but was %v", tos)
	}
}

func TestFuzzReproduction(t *testing.T) {

	code := []byte{
		SLOAD, 0, 0, 33,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "Index out of bounds" {
		t.Errorf("Expected tos to be 'Index out of bounds' error message but was %v", tos)
	}
}

func TestFuzzReproduction2(t *testing.T) {

	code := []byte{
		PUSH, 4, 46, 110, 66, 50, 255, SSTORE, 123, 119,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(false)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "Index out of bounds" {
		t.Errorf("Expected tos to be 'Index out of bounds' error message but was %v", tos)
	}
}

func TestFuzzReproduction3(t *testing.T) {

	code := []byte{
		29, 36, 222, 203, 229, 51, 52, 138, 144, 109, 48,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 300
	vm.Exec(true)

	tos, _ := vm.evaluationStack.Pop()

	if BigIntToString(tos) != "invalid data type supplied" {
		t.Errorf("Expected tos to be 'invalid data type supplied' error message but was %v", BigIntToString(tos))
	}
}

func TestFunctionCall(t *testing.T) {
	code := []byte{
		// start ABI
		DUP,
		PUSH, 0, 24,
		EQ,
		JMPIF, 24,
		DUP,
		PUSH, 0, 27,
		EQ,
		JMPIF, 27,
		HALT,
		// end ABI
		POP,
		SUB,
		HALT,
		POP,
		ADD,
		HALT,
	}

	vm := NewVM()
	vm.context.ContractAccount.Contract = code
	vm.context.MaxGasAmount = 50

	vm.context.TransactionData = []byte{
		PUSH, 0, 2,
		PUSH, 0, 5,
		PUSH, 0, 24, // Function hash
	}

	vm.Exec(true)
}
