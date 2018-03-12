package bazo_vm

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/sha3"
)

type VM struct {
	code            []int
	pc              int // Program counter
	evaluationStack Stack
}

func NewVM(startInstruction int) VM {
	return VM{
		code:            []int{},
		pc:              startInstruction,
		evaluationStack: NewStack(),
	}
}

// Private function, that can be activated by Exec call, useful for debugging
func (vm *VM) trace() {
	addr := vm.pc
	opCode := OpCodes[vm.code[vm.pc]]
	args := vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
	stack := vm.evaluationStack
	fmt.Printf("%04d: %s %v \t%v\n", addr, opCode.name, args, stack)
}

func (vm *VM) Exec(c []int, trace bool) {

	vm.code = c

	// Infinite Loop until return called
	for {
		if trace {
			vm.trace()
		}

		// Fetch
		opCode := vm.code[vm.pc]
		vm.pc++

		// Decode
		switch opCode {
		case PUSH:
			val := vm.code[vm.pc]
			vm.pc++
			vm.evaluationStack.Ipush(val)

		case ADD:
			right := vm.evaluationStack.Ipop()
			left := vm.evaluationStack.Ipop()
			vm.evaluationStack.Ipush(left + right)

		case SUB:
			right := vm.evaluationStack.Ipop()
			left := vm.evaluationStack.Ipop()
			vm.evaluationStack.Ipush(left - right)

		case MULT:
			right := vm.evaluationStack.Ipop()
			left := vm.evaluationStack.Ipop()
			vm.evaluationStack.Ipush(left * right)

		case DIV:
			right := vm.evaluationStack.Ipop()
			left := vm.evaluationStack.Ipop()
			vm.evaluationStack.Ipush(left / right)

		case MOD:
			right := vm.evaluationStack.Ipop()
			left := vm.evaluationStack.Ipop()
			vm.evaluationStack.Ipush(left % right)

		case EQ:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right == val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case NEQ:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right != val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case LT:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right < val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case GT:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right > val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case LTE:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right <= val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case GTE:
			val := vm.code[vm.pc]
			vm.pc++

			right, _ := vm.evaluationStack.Ipeek()

			if right >= val {
				vm.evaluationStack.Ipush(1)
			} else {
				vm.evaluationStack.Ipush(0)
			}

		case JMP:
			val := vm.code[vm.pc]

			//right := vm.evaluationStack.Ipop()

			vm.pc = val

		case JMPIF:
			val := vm.code[vm.pc]

			right := vm.evaluationStack.Ipop()

			if right == 1 {
				vm.pc = val
			} else {
				vm.pc++
			}

		case SHA3:
			val := vm.code[vm.pc]
			vm.pc++

			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, uint32(val))

			hasher := sha3.New256()
			hasher.Write(bs)
			vm.evaluationStack.Push(hasher.Sum(nil))

		case PRINT:
			val, _ := vm.evaluationStack.Ipeek()
			fmt.Println(val)

		case HALT:
			return
		}
	}
}
