package main

import "fmt"

type executionEngine struct {
	stack stack
}

func newExecutionEngine() executionEngine {
	return executionEngine{
		stack: newStack(),
	}
}

func (ee *executionEngine) executeOperation(opCode byte) {
	switch opCode {
	case PUSH1:
		ee.stack.push(int(opCode))
	case ADD:
		a, err1 := ee.stack.pop()
		b, err2 := ee.stack.pop()
		if err1 == nil && err2 == nil {
			ee.stack.push(a + b)
		} else {
			fmt.Println("Error")
		}
	case POP:
		ee.stack.pop()
	case SUB:
		a, err1 := ee.stack.pop()
		b, err2 := ee.stack.pop()
		if err1 == nil && err2 == nil {
			ee.stack.push(a - b)
		} else {
			fmt.Println("Error")
		}
	case MUL:
		//TODO
	}
}
