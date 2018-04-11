package bazo_vm

import (
	"fmt"
	"math/big"
	"reflect"

	"errors"

	"golang.org/x/crypto/sha3"
)

type VM struct {
	code            []byte
	pc              int // Program counter
	evaluationStack *Stack
	callStack       *CallStack
	context         *Context
}

func NewVM() VM {
	return VM{
		code:            []byte{},
		pc:              0,
		evaluationStack: NewStack(),
		callStack:       NewCallStack(),
		context:         NewContext(),
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

		if vm.pc+nargs < (len(vm.code) - vm.pc) {
			args = vm.code[vm.pc+2 : vm.pc+nargs+3]
			fmt.Printf("%04d: %-6s %-10v %v\n", addr, opCode.name, ByteArrayToInt(args), stack)
		}

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

func (vm *VM) Exec(trace bool) bool {

	vm.code = vm.context.contractAccount.Code

	if len(vm.code) > 100000 {
		vm.evaluationStack.Push(StrToBigInt("Instruction set to big"))
		return false
	}

	// Infinite Loop until return called
	for {
		if trace {
			vm.trace()
			//fmt.Println(vm.pc)
		}

		// Fetch
		opCode, err := vm.fetch()

		if err != nil {
			vm.evaluationStack.Push(StrToBigInt(err.Error()))
			return false
		}

		// Return false if instruction is not an opCode
		if len(OpCodes) < int(opCode) {
			vm.evaluationStack.Push(StrToBigInt("Not a valid opCode"))
			return false
		}

		// Substract gas used for operation
		if vm.context.maxGasAmount < OpCodes[int(opCode)].gasPrice {
			vm.evaluationStack.Push(StrToBigInt("out of gas"))
			return false
		} else {
			vm.context.maxGasAmount -= OpCodes[int(opCode)].gasPrice
		}

		// Decode
		switch opCode {

		case PUSH:
			arg, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			byteCount := int(arg) + 1 // Amount of bytes pushed

			var bigInt big.Int

			if int(byteCount)+1 >= (len(vm.code) - vm.pc) {
				vm.evaluationStack.Push(StrToBigInt("arguments exceeding instruction set"))
				return false
			}

			bigInt.SetBytes(vm.code[vm.pc : vm.pc+byteCount])

			vm.pc += byteCount //Sets the pc to the next opCode
			err = vm.evaluationStack.Push(bigInt)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case DUP:
			val, err := vm.evaluationStack.Peek()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			err = vm.evaluationStack.Push(val)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case ROLL:
			arg, err := vm.fetch() // arg shows how many have to be rolled

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			if int(arg) > vm.evaluationStack.GetLength() {
				vm.evaluationStack.Push(StrToBigInt("index out of bounds"))
				return false
			}

			index := vm.evaluationStack.GetLength() - (int(arg) + 2)

			newTos, err := vm.evaluationStack.PopIndexAt(index)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			err = vm.evaluationStack.Push(newTos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case ADD:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			left.Add(&left, &right)
			err := vm.evaluationStack.Push(left)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case SUB:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			left.Sub(&left, &right)
			err := vm.evaluationStack.Push(left)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case MULT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			left.Mul(&left, &right)
			err := vm.evaluationStack.Push(left)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case DIV:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			if right.Cmp(big.NewInt(0)) == 0 {
				vm.evaluationStack.Push(StrToBigInt("Division by Zero"))
				return false
			}

			left.Div(&left, &right)
			err := vm.evaluationStack.Push(left)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case MOD:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			if right.Cmp(big.NewInt(0)) == 0 {
				vm.evaluationStack.Push(StrToBigInt("Division by Zero"))
				return false
			}

			left.Mod(&left, &right)
			err := vm.evaluationStack.Push(left)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case NEG:
			tos, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos.Neg(&tos)

			vm.evaluationStack.Push(tos)

		case EQ:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case NEQ:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			if reflect.DeepEqual(left, right) {
				vm.evaluationStack.Push(*big.NewInt(0))
			} else {
				vm.evaluationStack.Push(*big.NewInt(1))
			}

		case LT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			value := left.Cmp(&right)

			if value == -1 {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case GT:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			value := left.Cmp(&right)

			if value == 1 {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case LTE:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			value := left.Cmp(&right)

			if value == -1 || value == 0 {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case GTE:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(rerr.Error()))
				return false
			}
			if lerr != nil {
				vm.evaluationStack.Push(StrToBigInt(lerr.Error()))
				return false
			}

			value := left.Cmp(&right)

			if value == 1 || value == 0 {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case SHIFTL:
			nrOfShifts, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos.Lsh(&tos, uint(nrOfShifts))
			err = vm.evaluationStack.Push(tos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case SHIFTR:
			nrOfShifts, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos.Rsh(&tos, uint(nrOfShifts))
			err = vm.evaluationStack.Push(tos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case NOP:
			_, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case JMP:
			nextInstruction, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			jumpTo := int(nextInstruction)
			vm.pc = jumpTo

		case JMPIF:
			nextInstruction, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			if right.Int64() == 1 {
				vm.pc = int(nextInstruction)
			}

		case CALL:
			jumpAddress, err := vm.fetch() // Shows where to jump after executing

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			if int(jumpAddress) == 0 || int(jumpAddress) > len(vm.code) {
				vm.evaluationStack.Push(StrToBigInt("JumpAddress out of bounds"))
				return false
			}

			argsToLoad, err := vm.fetch() // Shows how many elements have to be popped from evaluationStack

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			frame := &Frame{returnAddress: vm.pc, variables: make(map[int]big.Int)}

			for i := int(argsToLoad) - 1; i >= 0; i-- {
				frame.variables[i], err = vm.evaluationStack.Pop()
				if err != nil {
					vm.evaluationStack.Push(StrToBigInt(err.Error()))
					return false
				}
			}

			vm.callStack.Push(frame)
			vm.pc = int(jumpAddress) - 1

		case CALLEXT:
			transactionAddress := vm.code[vm.pc : vm.pc+32] // Addresses are 32 bytes
			vm.pc += 32                                     // Increase pc by address to get next instruction
			functionHash := vm.code[vm.pc : vm.pc+4]        // Function hash identifies function in external smart contract, first 4 byte of SHA3 hash
			vm.pc += 4                                      // Increase pc by function hash to get next instruction
			argsToLoad, err := vm.fetch()                   // Shows how many arguments to pop from stack and pass to external function

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			fmt.Println(transactionAddress, functionHash, argsToLoad)
			//TODO: Invoke new transaction with function hash and arguments, waiting for integration in bazo blockchain to finish

		case RET:
			callstackTos, err := vm.callStack.Peek()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.callStack.Pop()
			vm.pc = callstackTos.returnAddress

		case STORE:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.pc++
			address := vm.pc

			callstackTos, err := vm.callStack.Peek()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			callstackTos.variables[address] = right

		case LOAD:
			address, err := vm.fetch()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			callstackTos, err := vm.callStack.Peek()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			val := callstackTos.variables[int(address)]
			vm.evaluationStack.Push(val)

		case SHA3:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			hasher := sha3.New256()
			hasher.Write(right.Bytes())
			sha3_hash := hasher.Sum(nil)

			var bigInt big.Int
			bigInt.SetBytes(sha3_hash)

			err = vm.evaluationStack.Push(bigInt)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case ERRHALT:
			return false

		case HALT:
			return true
		}
	}
}

func (vm *VM) fetch() (element byte, err error) {
	tempPc := vm.pc
	vm.pc++
	if len(vm.code) > tempPc {
		return vm.code[tempPc], nil
	} else {
		return 0, errors.New("peek() on empty stack")
	}
}
