package bazo_vm

import (
	"fmt"
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
		case PUSH:
			byteCount := int(vm.code[vm.pc]) // Amount of bytes pushed
			vm.pc++                          //First byte

			var ba []byte
			for i := 0; i < byteCount; i++ {
				val := vm.code[vm.pc]
				ba = append(ba, val)
				vm.pc++
			}
			vm.evaluationStack.Push(ba)

		case PUSH1:
			val := vm.code[vm.pc]
			vm.pc++

			var byteArray []byte
			byteArray = append(byteArray, val)
			vm.evaluationStack.Push(byteArray)

		case PUSH2:
			var byteArray []byte
			for i := 0; i < 2; i++ {
				val := vm.code[vm.pc]
				byteArray = append(byteArray, val)
				vm.pc++

			}
			vm.evaluationStack.Push(byteArray)

		case PUSHS:
			val := vm.code[vm.pc]
			vm.pc++

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
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()
			result := ByteArrayToInt(left) + ByteArrayToInt(right)
			vm.evaluationStack.Push(IntToByteArray(result))

			/*
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
			*/
		case PRINT:
			val, _ := vm.evaluationStack.Peek()
			fmt.Println(val)

		case HALT:
			return
		}
	}
}
