package bazo_vm

const (
	PUSH = iota
	ADD
	SUB
	MULT
	DIV
	MOD
	EQ
	NEQ
	LT
	GT
	LTE
	GTE
	JMP
	JMPIF
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
	MOD:   OpCode{"mod", 0},
	EQ:    OpCode{"eq", 1},
	NEQ:   OpCode{"neq", 1},
	LT:    OpCode{"lt", 1},
	GT:    OpCode{"gt", 1},
	LTE:   OpCode{"lte", 1},
	GTE:   OpCode{"gte", 1},
	JMP:   OpCode{"jmp", 1},
	JMPIF: OpCode{"jmpif", 1},
	PRINT: OpCode{"print", 0},
	HALT:  OpCode{"halt", 0},
}
