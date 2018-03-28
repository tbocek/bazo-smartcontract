package bazo_vm

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
)

type VM struct {
	code            []byte
	pc              int // Program counter
	evaluationStack *Stack
	callStack       *CallStack
}

func NewVM(startInstruction int) VM {
	return VM{
		code:            []byte{},
		pc:              startInstruction,
		evaluationStack: NewStack(),
		callStack:       NewCallStack(),
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

	case "callext":
		nargs := int(vm.code[vm.pc+37])
		functionHash := vm.code[vm.pc+33 : vm.pc+37]
		address := vm.code[vm.pc+1 : vm.pc+33]
		fmt.Printf("%04d: %-6s %x %x %v %v\n", addr, opCode.name, address, functionHash, nargs, stack)

	default:
		args = vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
		fmt.Printf("%04d: %-6s %v %v\n", addr, opCode.name, args, stack)
	}
}

func (vm *VM) Exec(context Context, trace bool) bool {

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
				vm.evaluationStack.Push(StrToByteArray("out of gas"))
				return false
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

		case DUP:
			val, _ := vm.evaluationStack.Peek()
			vm.evaluationStack.Push(val)

		case ROLL:
			arg := vm.fetch() // arg shows how many have to be rolled
			newTos, err := vm.evaluationStack.PopIndexAt(vm.evaluationStack.GetLength() - (int(arg) + 2))

			if err != nil {
				return false
			}

			vm.evaluationStack.Push(newTos)

		case ADD:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			result := ByteArrayToInt(left) + ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case SUB:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			result := ByteArrayToInt(left) - ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case MULT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			result := ByteArrayToInt(left) * ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case DIV:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			result := ByteArrayToInt(left) / ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case MOD:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			result := ByteArrayToInt(left) % ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

		case EQ:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case NEQ:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(IntToByteArray(0))
			} else {
				vm.evaluationStack.Push(IntToByteArray(1))
			}

		case LT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if ByteArrayToInt(left) < ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case GT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if ByteArrayToInt(left) > ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case LTE:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if ByteArrayToInt(left) <= ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case GTE:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil || lerr != nil {
				return false
			}

			if ByteArrayToInt(left) >= ByteArrayToInt(right) {
				vm.evaluationStack.Push(IntToByteArray(1))
			} else {
				vm.evaluationStack.Push(IntToByteArray(0))
			}

		case SHIFTL:
			nrOfShifts := uint64(vm.fetch())

			ba, err := vm.evaluationStack.Pop()

			if err != nil {
				return false
			}

			value := ByteArrayToInt(ba)
			value = value << nrOfShifts
			vm.evaluationStack.Push(IntToByteArray(value))

		case SHIFTR:
			nrOfShifts := uint64(vm.fetch())

			ba, err := vm.evaluationStack.Pop()

			if err != nil {
				return false
			}

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
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				return false
			}

			if ByteArrayToInt(right) == 1 {
				vm.pc = val
			}

		case CALL:
			jumpAddress := int(vm.fetch()) // Shows where to jump after executing
			argsToLoad := int(vm.fetch())  // Shows how many elements have to be popped from evaluationStack

			frame := &Frame{returnAddress: vm.pc, variables: make(map[int][]byte)}

			var err error = nil
			for i := argsToLoad - 1; i >= 0; i-- {
				frame.variables[i], err = vm.evaluationStack.Pop()
				if err != nil {
					return false
				}
			}

			vm.callStack.Push(frame)
			vm.pc = jumpAddress - 1

		case CALLEXT:
			transactionAddress := vm.code[vm.pc : vm.pc+32] // Addresses are 32 bytes
			vm.pc += 32                                     // Increase pc by address to get next instruction
			functionHash := vm.code[vm.pc : vm.pc+4]        // Function hash identifies function in external smart contract, first 4 byte of SHA3 hash
			vm.pc += 4                                      // Increase pc by function hash to get next instruction
			argsToLoad := int(vm.fetch())                   // Shows how many arguments to pop from stack and pass to external function

			fmt.Println(transactionAddress, functionHash, argsToLoad)
			//TODO: Invoke new transaction with function hash and arguments, waiting for integration in bazo blockchain to finish

		case RET:
			returnAddress := vm.callStack.Peek().returnAddress
			vm.callStack.Pop()
			vm.pc = returnAddress

		case STORE:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				return false
			}

			vm.pc++
			address := vm.pc
			vm.callStack.Peek().variables[address] = right

		case LOAD:
			address := int(vm.fetch())

			val := vm.callStack.Peek().variables[address]
			vm.evaluationStack.Push(val)

		case SHA3:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				return false
			}

			hasher := sha3.New256()
			hasher.Write(right)

			sha3_hash := hasher.Sum(nil)
			vm.evaluationStack.Push(sha3_hash)

		case PRINT:
			val, _ := vm.evaluationStack.Peek()
			fmt.Println(val)

		case ERRHALT:
			return false

		case HALT:
			fmt.Println(vm.pc)
			return true
		}
	}
}

func (vm *VM) fetch() byte {
	tempPc := vm.pc
	vm.pc++
	return vm.code[tempPc]
}
