package bazo_vm

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
)

type VM struct {
	code            []byte
	pc              int // Program counter
	evaluationStack Stack
	callStack       CallStack
	memory          Memory
}

func NewVM(startInstruction int) VM {
	return VM{
		code:            []byte{},
		pc:              startInstruction,
		evaluationStack: NewStack(),
		callStack:       NewCallStack(),
		memory:          NewMemory(),
	}
}

// Private function, that can be activated by Exec call, useful for debugging
func (vm *VM) trace() {
	stack := vm.evaluationStack
	addr := vm.pc
	opCode := OpCodes[int(vm.code[vm.pc])]
	var args []byte

	switch opCode.name {
	case "push":
		nargs := int(vm.code[vm.pc+1])
		args = vm.code[vm.pc+2 : vm.pc+nargs+2]
		fmt.Printf("%04d: %-6s %-10v %v\n", addr, opCode.name, ByteArrayToInt(args), stack)

	case "pushs":
		tempPc := vm.pc
		arg := vm.code[tempPc]
		tempPc++

		firstRun := true
		var args []byte
		for firstRun == true || arg != 0x00 {
			firstRun = false
			arg = vm.code[tempPc]
			if arg != 0x00 {
				args = append(args, arg)
			}
			tempPc++
		}
		fmt.Printf("%04d: %-6s %-10v \t%v\n", addr, opCode.name, ByteArrayToString(args), stack)

	default:
		args = vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
		fmt.Printf("%04d: %-6s %v %v\n", addr, opCode.name, args, stack)
	}
}

func (vm *VM) Exec(context Context, trace bool) {

	vm.code = context.smartContract.data.code

	// Infinite Loop until return called
	for {

		if trace {
			vm.trace()
		}
		// Fetch
		opCode := vm.fetch()

		if opCode != HALT {
			if context.maxGasAmount <= 0 {
				return
			}
			context.maxGasAmount--
		}

		// Decode
		switch opCode {
		case PUSH:
			byteCount := int(vm.fetch()) // Amount of bytes pushed
			var ba byteArray = vm.code[vm.pc : vm.pc+byteCount]
			vm.pc += byteCount //Sets the pc to the next opCode
			vm.evaluationStack.Push(ba)

		case PUSHS:
			val := vm.fetch()

			firstRun := true
			var byteArray []byte
			for firstRun == true || val != 0x00 {
				firstRun = false
				byteArray = append(byteArray, val)
				val = vm.code[vm.pc]
				vm.pc++
			}
			vm.evaluationStack.Push(byteArray)

		case ADD:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) + ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case SUB:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) - ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case MULT:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) * ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case DIV:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) / ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case MOD:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) % ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case EQ:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case NEQ:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(IntToByteArray(0))
			} else {
				vm.evaluationStack.Push(IntToByteArray(1))
			}

		case LT:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if ByteArrayToInt(left) < ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case GT:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if ByteArrayToInt(left) > ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case LTE:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if ByteArrayToInt(left) <= ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case GTE:
			right, left := vm.evaluationStack.Pop(), vm.evaluationStack.Pop()

			if ByteArrayToInt(left) >= ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case SHIFTL:
			nrOfShifts := uint64(vm.fetch())

			ba := vm.evaluationStack.Pop()
			value := ByteArrayToInt(ba)
			value = value << nrOfShifts
			vm.evaluationStack.Push(IntToByteArray(value))

		case SHIFTR:
			nrOfShifts := uint64(vm.fetch())

			ba := vm.evaluationStack.Pop()
			value := ByteArrayToInt(ba)
			value = value >> nrOfShifts
			vm.evaluationStack.Push(IntToByteArray(value))

		case NOP:
			vm.fetch()

		case JMP:
			val := int(vm.fetch())
			vm.pc = val

		case JMPIF:
			val := int(vm.fetch())
			right := vm.evaluationStack.Pop()

			if ByteArrayToInt(right) == 1 {
				vm.pc = val
			}

		case CALL:
			jumpAddress := int(vm.fetch()) // Shows where to jump after executing
			argsToLoad := int(vm.fetch())  // Shows how many elements have to be popped from evaluationStack

			frame := Frame{returnAddress: vm.pc, variables: make(map[int][]byte)}

			for i := argsToLoad - 1; i >= 0; i-- {
				frame.variables[i] = vm.evaluationStack.Pop()
			}

			vm.callStack.Push(frame)
			vm.pc = jumpAddress - 1

		case RET:
			returnAddress := vm.callStack.Peek().returnAddress
			vm.callStack.Pop()
			vm.pc = returnAddress

		case STORE:
			right := vm.evaluationStack.Pop()
			vm.pc++
			address := vm.pc
			vm.callStack.Peek().variables[address] = right

		case LOAD:
			address := int(vm.fetch())

			val := vm.callStack.Peek().variables[address]
			vm.evaluationStack.Push(val)

		case MSTORE:
			right := vm.evaluationStack.Pop()
			vm.memory.Store(right)

		case MLOAD:
			index := int(vm.fetch())

			data, _ := vm.memory.Load(index)
			vm.evaluationStack.Push(data)

		case SHA3:
			right := vm.evaluationStack.Pop()

			hasher := sha3.New256()
			hasher.Write(right)

			sha3_hash := hasher.Sum(nil)
			vm.evaluationStack.Push(sha3_hash)

		case PRINT:
			val, _ := vm.evaluationStack.Peek()
			fmt.Println(val)

		case HALT:
			return
		}
	}
}

func (vm *VM) fetch() byte {
	tempPc := vm.pc
	vm.pc++
	return vm.code[tempPc]
}
