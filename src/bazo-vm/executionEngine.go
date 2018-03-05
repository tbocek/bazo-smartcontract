package main

import "fmt"

type executionEngine struct {
	stack stack
}

func newExecutionEngine() executionEngine{
	return executionEngine{
		stack: newStack(),
	}
}

func (ee executionEngine) executeOperation(opCode byte) {

	switch opCode {
	case PUSH0:
	case ADD:
		fmt.Println("ADD")
	case SUB:
		fmt.Println("SUP")
	case MUL:
		fmt.Println("MUL")
	}
}


