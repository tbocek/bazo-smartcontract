package bazo_vm

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
	"encoding/binary"
)

type VM struct {
	code            []instruction
	pc              int // Program counter
	evaluationStack Stack
}

func NewVM(startInstruction int) VM {
	return VM{
		code:            []instruction{},
		pc:              startInstruction,
		evaluationStack: NewStack(),
	}
}

// Private function, that can be activated by Exec call, useful for debugging
func (vm *VM) trace() {
	addr := vm.pc
	opCode := OpCodes[int(vm.code[vm.pc].opCode)]
	args := vm.code[vm.pc].args
	stack := vm.evaluationStack
	fmt.Printf("%04d: %s %v \t%v\n", addr, opCode.name, args, stack)
}

func (vm *VM) Exec(c []instruction, trace bool) {

	vm.code = c

	// Infinite Loop until return called
	for {

		if trace {
			vm.trace()
		}
		// Fetch
		opCode := vm.code[vm.pc].opCode
		args := vm.code[vm.pc].args

		vm.pc++

		// Decode
		switch opCode {
		case PUSHI:
			val := args
			vm.evaluationStack.Push(INT, val)

		case PUSHS:
			val := args
			vm.evaluationStack.Push(STRING, val)

		case ADD:

			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()
			vm.evaluationStack.PushInt(left + right)

		case SUB:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()
			vm.evaluationStack.PushInt(left - right)

		case MULT:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()
			vm.evaluationStack.PushInt(left * right)

		case DIV:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()
			vm.evaluationStack.PushInt(left / right)

		case MOD:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()
			vm.evaluationStack.PushInt(left % right)

		case EQ:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()

			if reflect.DeepEqual(left.byteArray, right.byteArray) {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case NEQ:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()

			if reflect.DeepEqual(left.byteArray, right.byteArray) {
				vm.evaluationStack.PushInt(0)
			} else {
				vm.evaluationStack.PushInt(1)
			}

		case LT:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()

			if left < right {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case GT:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()

			if left > right {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case LTE:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()

			if left <= right {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case GTE:
			right := vm.evaluationStack.PopInt()
			left := vm.evaluationStack.PopInt()

			if left >= right {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case SHIFTL:
			var nrOfShifts uint64 = binary.LittleEndian.Uint64(args)
			ba := vm.evaluationStack.Pop().byteArray
			value := ByteArrayToInt(ba)
			value = value << nrOfShifts
			vm.evaluationStack.Push(INT, IntToByteArray(value))

		case SHIFTR:
			var nrOfShifts uint64 = binary.LittleEndian.Uint64(args)
			ba := vm.evaluationStack.Pop().byteArray

			value := ByteArrayToInt(ba)
			value = value >> nrOfShifts
			vm.evaluationStack.Push(INT, IntToByteArray(value))


		case JMP:
			val := ByteArrayToInt(args)
			vm.pc = val

		case JMPIF:
			val := ByteArrayToInt(args)
			right := vm.evaluationStack.PopInt()

			if right == 1 {
				vm.pc = val
			} else {
				vm.pc++
			}

		case SHA3:
			right := vm.evaluationStack.Pop()

			hasher := sha3.New256()
			hasher.Write(right.byteArray)

			sha3_hash := hex.EncodeToString(hasher.Sum(nil))
			vm.evaluationStack.PushStr(sha3_hash)

		case PRINT:
			val, _ := vm.evaluationStack.PeekInt()
			fmt.Println(val)

		case HALT:
			return
		}
	}
}
