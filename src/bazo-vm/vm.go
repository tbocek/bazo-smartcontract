package bazo_vm

import (
	"fmt"
	"math/big"
	"reflect"
	"errors"
	"crypto/ecdsa"
	"crypto/elliptic"
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

		//TODO - Fix CALLEXT case, leads to index out of bounds exception
	/*case "callext":
	address := vm.code[vm.pc+1 : vm.pc+33]
	functionHash := vm.code[vm.pc+33 : vm.pc+37]
	nargs := int(vm.code[vm.pc+37])

	fmt.Printf("%04d: %-6s %x %x %v %v\n", addr, opCode.name, address, functionHash, nargs, stack)
	*/

	case "mappush":
	case "mapgetval":
	case "newarr":
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

	vm.code = vm.context.ContractAccount.Contract

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
		if vm.context.MaxGasAmount < OpCodes[int(opCode)].gasPrice {
			vm.evaluationStack.Push(StrToBigInt("out of gas"))
			return false
		} else {
			vm.context.MaxGasAmount -= OpCodes[int(opCode)].gasPrice
		}

		// Decode
		switch opCode {

		case PUSH:
			arg, errArg1 := vm.fetch()
			byteCount := int(arg) + 1 // Amount of bytes pushed
			bytes, errArg2 := vm.fetchMany(byteCount)

			if !vm.checkErrors([]error{errArg1, errArg2}) {
				return false
			}

			var bigInt big.Int
			bigInt.SetBytes(bytes)

			err = vm.evaluationStack.Push(bigInt)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case DUP:
			val, err := vm.evaluationStack.Peek()

			if !vm.checkErrors([]error{err}) {
				return false
			}

			err = vm.evaluationStack.Push(val)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case ROLL:
			arg, err := vm.fetch() // arg shows how many have to be rolled
			index := vm.evaluationStack.GetLength() - (int(arg) + 2)

			if !vm.checkErrors([]error{err}) {
				return false
			}

			if index != -1 {
				if int(arg) >= vm.evaluationStack.GetLength() {
					vm.evaluationStack.Push(StrToBigInt("index out of bounds"))
					return false
				}

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
			}

		case ADD:
			right, rerr := vm.evaluationStack.Pop()
			left, lerr := vm.evaluationStack.Pop()

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
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

			if !vm.checkErrors([]error{rerr, lerr}) {
				return false
			}

			value := left.Cmp(&right)

			if value == 1 || value == 0 {
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

		case SHIFTL:
			nrOfShifts, errArg := vm.fetch()
			tos, errStack := vm.evaluationStack.Pop()

			if !vm.checkErrors([]error{errArg, errStack}) {
				return false
			}

			tos.Lsh(&tos, uint(nrOfShifts))
			err = vm.evaluationStack.Push(tos)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case SHIFTR:
			nrOfShifts, errArg := vm.fetch()
			tos, errStack := vm.evaluationStack.Pop()

			if !vm.checkErrors([]error{errArg, errStack}) {
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
			nextInstruction, errArg := vm.fetch()
			right, errStack := vm.evaluationStack.Pop()

			if !vm.checkErrors([]error{errArg, errStack}) {
				return false
			}

			if right.Int64() == 1 {
				vm.pc = int(nextInstruction)
			}

		case CALL:
			jumpAddress, errArg1 := vm.fetch() // Shows where to jump after executing
			argsToLoad, errArg2 := vm.fetch()  // Shows how many elements have to be popped from evaluationStack

			if !vm.checkErrors([]error{errArg1, errArg2}) {
				return false
			}

			if int(jumpAddress) == 0 || int(jumpAddress) > len(vm.code) {
				vm.evaluationStack.Push(StrToBigInt("JumpAddress out of bounds"))
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
			transactionAddress, errArg1 := vm.fetchMany(32) // Addresses are 32 bytes (var name: transactionAddress)
			functionHash, errArg2 := vm.fetchMany(4)        // Function hash identifies function in external smart contract, first 4 byte of SHA3 hash (var name: functionHash)
			argsToLoad, errArg3 := vm.fetch()               // Shows how many arguments to pop from stack and pass to external function (var name: argsToLoad)

			if !vm.checkErrors([]error{errArg1, errArg2, errArg3}) {
				return false
			}

			fmt.Sprint("CALLEXT", transactionAddress, functionHash, argsToLoad)
			//TODO: Invoke new transaction with function hash and arguments, waiting for integration in bazo blockchain to finish

		case RET:
			callstackTos, err := vm.callStack.Peek()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			vm.callStack.Pop()
			vm.pc = callstackTos.returnAddress

		case SIZE:
			right, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			err = vm.evaluationStack.Push(*big.NewInt(int64(getElementMemoryUsage(right.BitLen()))))

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case SSTORE:
			index, err := vm.fetch()

			if !vm.checkErrors([]error{err}) {
				return false
			}

			value, err := vm.evaluationStack.Pop()

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

			if len(vm.context.ContractAccount.ContractVariables) <= int(index) {
				vm.evaluationStack.Push(StrToBigInt("Index out of bounds"))
				return false
			}

			vm.context.ContractAccount.ContractVariables[int(index)] = value

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

		case SLOAD:
			const HASHLENGTH = 1
			index, err := vm.fetchMany(HASHLENGTH)

			if !vm.checkErrors([]error{err}) {
				return false
			}

			if len(vm.context.ContractAccount.ContractVariables) <= ByteArrayToInt(index) {
				vm.evaluationStack.Push(StrToBigInt("Index out of bounds"))
				return false
			}

			value := vm.context.ContractAccount.ContractVariables[ByteArrayToInt(index)]

			err = vm.evaluationStack.Push(value)

			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(err.Error()))
				return false
			}

		case LOAD:
			address, errArg := vm.fetch()
			callstackTos, errCallStack := vm.callStack.Peek()

			if !vm.checkErrors([]error{errArg, errCallStack}) {
				return false
			}

			val := callstackTos.variables[int(address)]
			vm.evaluationStack.Push(val)

		case NEWMAP:
			//TODO guards and errors
			datastructure, _ := vm.fetch()

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
			a := NewArray()
			vm.evaluationStack.Push(a.ToBigInt())

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

			arr := Array(a.Bytes())
			err := arr.Append(v)
			if err != nil {
				vm.evaluationStack.Push(StrToBigInt("Invalid argument size of ARRAPPEND"))
				return false
			}


			vm.evaluationStack.Push(arr.ToBigInt())

		case ARRREMOVE:
			a, aerr := vm.evaluationStack.Pop()
			index := BaToUI16(vm.code[vm.pc : vm.pc+2])
			vm.pc += 2
			if aerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			arr, perr := ArrayFromBigInt(a)
			if perr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			rerr := arr.Remove(index)
			if rerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			derr := arr.DecrementSize()
			if derr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			vm.evaluationStack.Push(arr.ToBigInt())

		case ARRAT:
			a, aerr := vm.evaluationStack.Peek()
			index := BaToUI16(vm.code[vm.pc : vm.pc+2])
			vm.pc += 2

			if aerr != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			arr, err := ArrayFromBigInt(a)
			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}

			e, err := arr.At(index)
			if err != nil {
				vm.evaluationStack.Push(StrToBigInt(aerr.Error()))
				return false
			}
			result := big.Int{}
			result.SetBytes(e)
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

		case CHECKSIG:
			publicKeySig, errArg1 := vm.evaluationStack.Pop() // PubKeySig
			hash, errArg2 := vm.evaluationStack.Pop()         // Hash

			if !vm.checkErrors([]error{errArg1, errArg2}) {
				return false
			}

			if len(publicKeySig.Bytes()) != 64 {
				vm.evaluationStack.Push(StrToBigInt("Not a valid address"))
				return false
			}

			if len(hash.Bytes()) != 32 {
				vm.evaluationStack.Push(StrToBigInt("Not a valid hash"))
				return false
			}

			pubKey1Sig1, pubKey2Sig1 := new(big.Int), new(big.Int)
			r, s := new(big.Int), new(big.Int)

			pubKey1Sig1.SetBytes(publicKeySig.Bytes()[:32])
			pubKey2Sig1.SetBytes(publicKeySig.Bytes()[32:])

			r.SetBytes(vm.context.ContractTx.Sig1[:32])
			s.SetBytes(vm.context.ContractTx.Sig1[32:])

			pubKey := ecdsa.PublicKey{elliptic.P256(), pubKey1Sig1, pubKey2Sig1}

			if ecdsa.Verify(&pubKey, hash.Bytes(), r, s) {
				fmt.Println("Valid Sig", pubKey, hash.Bytes())
				vm.evaluationStack.Push(*big.NewInt(1))
			} else {
				vm.evaluationStack.Push(*big.NewInt(0))
			}

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

func (vm *VM) fetch() (element byte, err error) {
	tempPc := vm.pc
	if len(vm.code) > tempPc {
		vm.pc++
		return vm.code[tempPc], nil
	} else {
		return 0, errors.New("instructionSet out of bounds")
	}
}

func (vm *VM) fetchMany(argument int) (elements []byte, err error) {
	tempPc := vm.pc
	if len(vm.code)-tempPc > argument {
		vm.pc += argument
		return vm.code[tempPc : tempPc+argument], nil
	} else {
		return []byte{}, errors.New("instructionSet out of bounds")
	}
}

func (vm *VM) checkErrors(errors []error) bool {
	for i, err := range errors {
		if err != nil {
			vm.evaluationStack.Push(StrToBigInt(errors[i].Error()))
			return false
		}
	}
	return true
}
