package main

import "fmt"

type VM struct {
	code            []int
	pc              int // Program counter
	evaluationStack Stack
}

func NewVM() VM {
	return VM{
		code:            []int{},
		pc:              0,
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

	// Infinite Loop until break called
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
			vm.evaluationStack.Push(val)

		case ADD:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()
			vm.evaluationStack.Push(left + right)

		case SUB:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()
			vm.evaluationStack.Push(left - right)

		case MULT:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()
			vm.evaluationStack.Push(left * right)

		case DIV:
			right := vm.evaluationStack.Pop()
			left := vm.evaluationStack.Pop()
			vm.evaluationStack.Push(left / right)

		case PRINT:
			val, _ := vm.evaluationStack.Peek()
			fmt.Println(val)

		case HALT:
			return
		}
	}
}
