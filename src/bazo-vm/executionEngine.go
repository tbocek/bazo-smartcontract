package main

import "fmt"

type executionEngine struct {
	stack stack
}

func (ee executionEngine) executeOperation(opCode []byte) {

	switch opCode {
	case PUSH0:
		stack.push(opCode)
	case ADD:
		fmt.Println("ADD")
	case SUB:
		fmt.Println("SUP")
	case MUL:
		fmt.Println("MUL")
	}
}


