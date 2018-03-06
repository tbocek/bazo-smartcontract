package main

const (
	PUSH = iota
	ADD
	SUB
	MULT
	DIV
	PRINT
	HALT
)

type OpCode struct {
	name  string
	nargs int
}

var OpCodes = map[int]OpCode{
	PUSH:  OpCode{"push", 1},
	ADD:   OpCode{"add", 0},
	SUB:   OpCode{"sub", 0},
	MULT:  OpCode{"mult", 0},
	DIV:   OpCode{"div", 0},
	PRINT: OpCode{"print", 0},
	HALT:  OpCode{"halt", 0},
}
