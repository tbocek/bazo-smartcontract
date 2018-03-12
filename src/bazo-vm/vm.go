package bazo_vm

import (
	"encoding/binary"
	"fmt"

	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type VM struct {
	code            []byte
	pc              int // Program counter
	evaluationStack Stack
}

func NewVM(startInstruction int) VM {
	return VM{
		code:            []byte{},
		pc:              startInstruction,
		evaluationStack: NewStack(),
	}
}

// Private function, that can be activated by Exec call, useful for debugging
func (vm *VM) trace() {
	addr := vm.pc
	opCode := OpCodes[int(vm.code[vm.pc])]
	args := vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
	stack := vm.evaluationStack
	fmt.Printf("%04d: %s %v \t%v\n", addr, opCode.name, args, stack)
}

func (vm *VM) Exec(c []byte, trace bool) {

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
		case PUSHI:
			val := int(vm.code[vm.pc])
			vm.pc++
			vm.evaluationStack.PushInt(val)

		case PUSHS:
			val := vm.code[vm.pc]
			vm.pc++

			firstRun := true
			word := ""
			for firstRun == true || val != 0x00 {
				firstRun = false
				word += string([]byte{val})
				val = vm.code[vm.pc]
				vm.pc++
			}
			vm.evaluationStack.PushStr(word)

		case ADD:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()

			if left.dataType == STRING && right.dataType == STRING {
				result := string(left.byteArray) + string(right.byteArray)
				vm.evaluationStack.PushStr(string(result))
			}

			if left.dataType == INT && right.dataType == INT {
				vm.evaluationStack.PushInt(int(left.byteArray[0]) + int(right.byteArray[0]))
			}

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
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right == val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case NEQ:
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right != val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case LT:
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right < val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case GT:
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right > val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case LTE:
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right <= val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case GTE:
			val := int(vm.code[vm.pc])
			vm.pc++

			right, _ := vm.evaluationStack.PeekInt()

			if right >= val {
				vm.evaluationStack.PushInt(1)
			} else {
				vm.evaluationStack.PushInt(0)
			}

		case JMP:
			val := int(vm.code[vm.pc])
			vm.pc = val

		case JMPIF:
			val := int(vm.code[vm.pc])

			right := vm.evaluationStack.PopInt()

			if right == 1 {
				vm.pc = val
			} else {
				vm.pc++
			}

		case SHA3:
			val := int(vm.code[vm.pc])
			vm.pc++

			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, uint32(val))

			hasher := sha3.New256()
			hasher.Write(bs)

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
