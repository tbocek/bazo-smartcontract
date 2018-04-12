package bazo_vm

import (
	"fmt"
	"math/big"
	"reflect"
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

	case "mappush":
	case "mapgetval":
	case "arrappend":
	case "arrremove":
	case "arrat":
		args = vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
		fmt.Printf("%04d: %-6s %v ", addr, opCode.name, args)

		for _, e := range stack.Stack {
			fmt.Printf("%# x", e.Bytes())
			fmt.Printf("\n")
		}

		fmt.Printf("\n")

	default:
		args = vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
		fmt.Printf("%04d: %-6s %v %v\n", addr, opCode.name, args, stack)
	}
}

func (vm *VM) Exec(trace bool) bool {

	vm.code = vm.context.contractAccount.Code

	// Infinite Loop until return called
	for {
		if trace {
			vm.trace()
		}

		// Fetch
		opCode := vm.fetch()

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
			byteCount := int(vm.fetch()) + 1 // Amount of bytes pushed
			var bigInt big.Int

			if byteCount >= (len(vm.code) - vm.pc) {
				vm.evaluationStack.Push(StrToBigInt("arguments exceeding instruction set"))
				return false
			}

			bigInt.SetBytes(vm.code[vm.pc : vm.pc+byteCount])

			vm.pc += byteCount //Sets the pc to the next opCode
			err := vm.evaluationStack.Push(bigInt)

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
			arg := vm.fetch() // arg shows how many have to be rolled
			newTos, err := vm.evaluationStack.PopIndexAt(vm.evaluationStack.GetLength() - (int(arg) + 2))

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.evaluationStack.Push(newTos)

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
			nrOfShifts := uint(vm.fetch())

			tos, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos.Lsh(&tos, nrOfShifts)
			err = vm.evaluationStack.Push(tos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case SHIFTR:
			nrOfShifts := uint(vm.fetch())

			tos, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			tos.Rsh(&tos, nrOfShifts)
			err = vm.evaluationStack.Push(tos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case NOP:
			vm.fetch()

		case JMP:
			val := int(vm.fetch())
			vm.pc = val

		case JMPIF:
			val := int(vm.fetch())
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			if right.Int64() == 1 {
				vm.pc = val
			}

		case CALL:
			jumpAddress := int(vm.fetch()) // Shows where to jump after executing
			argsToLoad := int(vm.fetch())  // Shows how many elements have to be popped from evaluationStack

			frame := &Frame{returnAddress: vm.pc, variables: make(map[int]big.Int)}

			var err error = nil
			for i := argsToLoad - 1; i >= 0; i-- {
				frame.variables[i], err = vm.evaluationStack.Pop()
				if err != nil {
					vm.evaluationStack.Push(StrToBigInt(err.Error()))
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
			argsToLoad := int(vm.fetch())                   // Shows how many arguments to pop from Stack and pass to external function

			fmt.Println(transactionAddress, functionHash, argsToLoad)
			//TODO: Invoke new transaction with function hash and arguments, waiting for integration in bazo blockchain to finish

		case RET:
			returnAddress := vm.callStack.Peek().returnAddress
			vm.callStack.Pop()
			vm.pc = returnAddress

		case SSTORE:
			key := vm.code[vm.pc : vm.pc+1]
			vm.pc += 1

			value, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.context.contractAccount.ContractVariables[ByteArrayToInt(key)] = value


		case STORE:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.pc++
			address := vm.pc
			vm.callStack.Peek().variables[address] = right

		case SLOAD:
			const HASHLENGTH = 1
			key := vm.code[vm.pc : vm.pc+HASHLENGTH]
			vm.pc += HASHLENGTH

			value := vm.context.contractAccount.ContractVariables[ByteArrayToInt(key)]
			vm.evaluationStack.Push(value);


		case LOAD:
			address := int(vm.fetch())

			val := vm.callStack.Peek().variables[address]
			vm.evaluationStack.Push(val)

		case NEWMAP:
			datastructure := vm.fetch()
			ba := []byte{datastructure}

			keylength := vm.code[vm.pc : vm.pc+8]
			ba = append(ba, keylength...)
			vm.pc += 8

			valuelength := vm.code[vm.pc : vm.pc+8]
			ba = append(ba, valuelength...)
			vm.pc += 8

			size := vm.code[vm.pc : vm.pc+8]
			ba = append(ba, size...)
			vm.pc += 8

			m := big.Int{}
			m.SetBytes(ba)
			vm.evaluationStack.Push(m)

		case MAPPUSH:
			k, kerr := vm.evaluationStack.Pop()
			v, verr := vm.evaluationStack.Pop()
			m, merr := vm.evaluationStack.Pop()


			if kerr != nil {
				vm.evaluationStack.Push(StrToBigInt(kerr.Error()))
				return false
			}
			if verr != nil {
				vm.evaluationStack.Push(StrToBigInt(verr.Error()))
				return false
			}
			if merr != nil {
				vm.evaluationStack.Push(StrToBigInt(merr.Error()))
				return false
			}

			mba := m.Bytes()


			mba = append(mba, k.Bytes()...)
			mba = append(mba, v.Bytes()...)

			s := BaToi(mba[17:25])
			s++
			sba := IToBA(s)

			for i, b := range sba {
				mba[17+i] = b
			}

			m.SetBytes(mba)
			vm.evaluationStack.Push(m)

		case MAPGETVAL:
			tmpKey, kerr := vm.evaluationStack.Pop()
			m, merr := vm.evaluationStack.Pop()

			if kerr != nil {
				vm.evaluationStack.Push(StrToBigInt(kerr.Error()))
				return false
			}

			if merr != nil {
				vm.evaluationStack.Push(StrToBigInt(merr.Error()))
				return false
			}

			key := tmpKey.Bytes()
			mba := m.Bytes()

			kl, vl, size := getMapProperties(mba)
			mel := kl + vl

			data := mba[25:]
			var i uint64 = 0
			var offset uint64 = 0
			for ; i < size; i++ {
				bv := offset + kl
				k := data[offset:bv]
				if reflect.DeepEqual(k, key) {
					v := data[bv:bv+vl]
					r := big.Int{}
					r.SetBytes(v)
					vm.evaluationStack.Push(r)
				}
				offset += mel
			}

		case NEWARR:
			ba := []byte{0x002,}
			valuelength := vm.code[vm.pc : vm.pc+8]
			ba = append(ba, valuelength...)
			vm.pc += 8

			size := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}
			ba = append(ba, size...)

			arr := big.Int{}
			arr.SetBytes(ba)
			vm.evaluationStack.Push(arr)

		case ARRAPPEND:
			v, verr := vm.evaluationStack.Pop()
			a, aerr := vm.evaluationStack.Pop()

			if aerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			if verr != nil {
				vm.evaluationStack.Push(StrToBigInt(verr.Error()))
				return false
			}

			ba := a.Bytes()
			vb := v.Bytes()
			size := BaToi(ba[9:17])
			size++
			newSize := IToBA(size)

			for i, e := range newSize {
				ba[9+i] = e
			}

			//TODO unwanted cast
			if uint64(len(vb)) != BaToi(ba[1:9]) {
				vm.evaluationStack.Push(StrToBigInt("Invalid argument size of ARRAPPEND"))
				return false
			}


			ba = append(ba, vb...)

			arr := big.Int{}
			arr.SetBytes(ba)
			vm.evaluationStack.Push(arr)

		case ARRREMOVE:
			a, aerr := vm.evaluationStack.Pop()
			index := BaToi(vm.code[vm.pc : vm.pc+8])
			//TODO Size check
			//TODO Datastructure check
			vm.pc += 8

			if aerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			arr := a.Bytes()
			elementByteSize := BaToi(arr[1:9])
			offset := uint64(1 + 8 + 8)
			elementToRemove := offset + (elementByteSize * index)

			size := BaToi(arr[9:17])
			size--
			newSize := IToBA(size)

			for i, e := range newSize {
				arr[9+i] = e
			}

			left := arr[:elementToRemove]
			right := arr[elementToRemove + elementByteSize : ]

			arr = left
			arr = append(arr, right...)
			result := big.Int{}
			result.SetBytes(arr)
			vm.evaluationStack.Push(result)

		case ARRAT:
			a, aerr := vm.evaluationStack.Pop()
			index := BaToi(vm.code[vm.pc : vm.pc+8])
			//TODO Size check
			//TODO Datastructure check
			vm.pc += 8

			if aerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			arr := a.Bytes()
			elementByteSize := BaToi(arr[1:9])
			offset := uint64(1 + 8 + 8)
			elementStart := offset + (elementByteSize * index)

			element := arr[elementStart : elementStart + elementByteSize]
			result := big.Int{}
			result.SetBytes(element)
			vm.evaluationStack.Push(result)

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

		case PRINT:
			val, _ := vm.evaluationStack.Peek()
			fmt.Println(val)

		case ERRHALT:
			return false

		case HALT:
			return true
		}
	}
}
func getMapProperties(mba []byte) (uint64, uint64, uint64) {
	kl := BaToi(mba[1:9])
	vl := BaToi(mba[9:17])
	size := BaToi(mba[17:25])
	return kl, vl, size
}

func (vm *VM) fetch() byte {
	tempPc := vm.pc
	vm.pc++
	return vm.code[tempPc]
}
